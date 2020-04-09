package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/phelmkamp/magnum/gen"
)

const (
	accessTemplate = "%s.%s"
)

var (
	goFileRegEx  = regexp.MustCompile(`.+\.go$`)
	enumTagRegEx = regexp.MustCompile(`enum:".+"`)
)

func initFile(origPath string) *os.File {
	filename := strings.Replace(origPath, ".go", "_enum.go", 1)
	log.Printf("Creating file: %s\n", filename)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create() failed: %v\n", err)
	}
	return f
}

func first(s string) (string, int) {
	if s == "" {
		return "", 0
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(r), n
}

func upperFirst(s string) string {
	f, n := first(s)
	return strings.ToUpper(f) + s[n:]
}

func main() {
	var root string
	flag.StringVar(&root, "path", ".", "directory path to scan for *.go files")
	flag.Parse()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && goFileRegEx.MatchString(info.Name()) {
			cleanPath := filepath.Clean(path)

			log.Printf("Parsing file: %s\n", cleanPath)
			fset := token.NewFileSet()
			astFile, err := parser.ParseFile(fset, cleanPath, nil, 0)
			if err != nil {
				return fmt.Errorf("parser.ParseFile() failed: %w", err)
			}

			genFile := gen.NewFile(astFile.Name.Name)
			importPaths := make([]string, 0, 8)

			ast.Inspect(astFile, func(n ast.Node) bool {
				var expr ast.Expr
				var typeName string
				switch nt := n.(type) {
				case *ast.TypeSpec:
					expr = nt.Type
					typeName = nt.Name.Name
				case *ast.ImportSpec:
					p := strings.Trim(nt.Path.Value, `"`)
					importPaths = append(importPaths, p)
				}

				if expr == nil {
					return true
				}

				st, ok := expr.(*ast.StructType)
				if !ok {
					return true
				}

				log.Printf("Found struct: %s\n", typeName)

				rcvName, _ := first(typeName)
				rcvName = strings.ToLower(rcvName)

				valueFns := make(gen.Funcs, 0, 10)
				getters := make(gen.Funcs, 0, 10)

				for _, f := range st.Fields.List {
					if f.Tag == nil {
						continue
					}

					enumTag := enumTagRegEx.FindString(f.Tag.Value)
					if enumTag == "" {
						continue
					}

					log.Printf("Found enum tag: %s\n", enumTag)
					enumTag = strings.TrimPrefix(enumTag, "enum:\"")
					enumTag = strings.TrimSuffix(enumTag, "\"")

					var fldPkg, fldType string
					switch ft := f.Type.(type) {
					case *ast.Ident:
						fldType = ft.Name
					case *ast.SelectorExpr:
						// package.type
						fldPkg = ft.X.(*ast.Ident).Name
						fldType = fmt.Sprintf(accessTemplate, fldPkg, ft.Sel.Name)
					case *ast.ArrayType:
						switch elt := ft.Elt.(type) {
						case *ast.Ident:
							fldType = "[]" + elt.Name
						case *ast.SelectorExpr:
							// package.type
							fldPkg = elt.X.(*ast.Ident).Name
							fldType = fmt.Sprintf(accessTemplate, "[]"+fldPkg, elt.Sel.Name)
						}
					case *ast.MapType:
						fldType = fmt.Sprintf("map[%s]%s", ft.Key.(*ast.Ident).Name, ft.Value.(*ast.Ident).Name)
					default:
						log.Printf("Unsupported field type: %v\n", ft)
						continue
					}

					if fldPkg != "" {
						var importPath string
						for _, s := range importPaths {
							subs := strings.Split(s, "/")
							last := subs[len(subs)-1]
							if last == fldPkg {
								importPath = s
								break
							}
						}
						log.Printf("Adding import: \"%s\"\n", importPath)
						genFile.Imports[importPath] = struct{}{}
					}

					fldName := f.Names[0].Name
					var valsBldr *strings.Builder
					if fldName == "name" {
						valsBldr = &strings.Builder{}
						fn := gen.Func{
							Name:    upperFirst(typeName) + "s",
							RetVals: "[]" + typeName,
							Misc: map[string]interface{}{
								"Values": valsBldr,
							},
							Tmpl: "values",
						}
						genFile.Funcs = append(genFile.Funcs, fn)
						log.Printf("Adding function: \"%s\"\n", fn.Name)
					}

					vals := strings.Split(enumTag, ",")
					for i, v := range vals {
						v = strings.TrimSpace(v)

						var sb *strings.Builder
						if i >= len(valueFns) {
							sb = &strings.Builder{}
							fn := gen.Func{
								Name:    upperFirst(v),
								RetVals: typeName,
								Misc: map[string]interface{}{
									"Fields": sb,
								},
								Tmpl: "value",
							}
							valueFns = append(valueFns, fn)
							log.Printf("Adding function: \"%s\"\n", fn.Name)
						} else {
							sb = valueFns[i].Misc["Fields"].(*strings.Builder)
							sb.WriteString("\n\t\t")
						}

						sb.WriteString(fldName)
						sb.WriteString(": ")
						if fldPkg != "" {
							sb.WriteString(fldPkg)
							sb.WriteString(".")
						}
						if fldType == "string" {
							sb.WriteString(`"`)
							sb.WriteString(v)
							sb.WriteString(`"`)
						} else {
							sb.WriteString(v)
						}
						sb.WriteString(",")

						if fldName == "name" {
							if valsBldr.Len() > 0 {
								valsBldr.WriteString(", ")
							}
							valsBldr.WriteString(upperFirst(v))
							valsBldr.WriteString("()")
							valueFns[i].Misc["Value"] = `"` + v + `"`
						}
					}

					fn := gen.Func{
						RcvName: rcvName,
						RcvType: typeName,
						RetVals: fldType,
						Misc: map[string]interface{}{
							"FldName": fldName,
						},
						Tmpl: "getter",
					}
					if fldName == "name" {
						fn.Name = "String"
					} else {
						fn.Name = upperFirst(fldName)
					}
					getters = append(getters, fn)
					log.Printf("Adding method: \"%s.%s\"\n", typeName, fn.Name)
				}

				genFile.Funcs = append(genFile.Funcs, valueFns...)
				genFile.Funcs = append(genFile.Funcs, getters...)

				return true
			})

			if len(genFile.Funcs) < 1 {
				return nil
			}

			osFile := initFile(cleanPath)
			defer func() {
				if err := osFile.Close(); err != nil {
					log.Printf("File.Close() failed: %v\n", err)
				}
			}()

			if _, err := osFile.WriteString(genFile.String()); err != nil {
				log.Fatalf("File.WriteString() failed: %v\n", err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
