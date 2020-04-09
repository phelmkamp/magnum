package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

var _getter_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x28\x4a\x2d\x29\x2d\xca\x2b\x56\x28\xc9\x48\x05\x89\x06\x25\x97\x85\x54\x16\xa4\xd6\xd6\xaa\x17\x83\xb8\xbe\x99\xc5\xc9\x7a\x6e\x39\x29\x10\xc5\x7a\xbc\x5c\x69\xa5\x79\xc9\x0a\x1a\x10\x85\x50\x13\x90\x75\x69\x22\x4c\xd6\x00\xb3\x83\x52\x4b\xc2\x12\x73\x8a\x41\xca\x78\xb9\x38\x21\x96\x29\x20\x6b\xd7\xc3\xb0\x85\x97\xab\x16\x10\x00\x00\xff\xff\xd7\xe8\xa7\xf6\xa2\x00\x00\x00")

func getter_tmpl() ([]byte, error) {
	return bindata_read(
		_getter_tmpl,
		"getter.tmpl",
	)
}

var _value_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x28\x4a\x2d\x29\x2d\xca\x2b\x56\x28\xc9\x48\x05\x89\xfa\x66\x16\x27\xeb\x85\x25\xe6\x94\x82\xe4\xaa\xab\xf5\x82\x52\x4b\xc2\x12\x73\x8a\x6b\x6b\xf5\x78\xb9\xd2\x4a\xf3\x92\x11\x3a\x35\x34\x51\xe4\x15\xaa\x79\xb9\x38\x21\x86\xa1\x88\x57\x57\x2b\x55\x2b\xd5\xd6\xf2\x72\x71\x72\xc2\x8c\x77\xcb\x4c\xcd\x49\x29\x06\x8b\x55\x57\x2b\xd5\x82\x65\x6b\x01\x01\x00\x00\xff\xff\x5e\x12\x76\x20\x95\x00\x00\x00")

func value_tmpl() ([]byte, error) {
	return bindata_read(
		_value_tmpl,
		"value.tmpl",
	)
}

var _values_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x28\x4a\x2d\x29\x2d\xca\x2b\x56\x48\xcc\xc9\x51\x28\xc8\x2f\x2e\xce\x4c\xca\x49\x45\x48\xeb\xf1\x72\xa5\x95\xe6\x25\x23\x04\x34\x34\x41\xec\xa0\xd4\x92\xb0\xc4\x9c\xe2\xda\x5a\x85\x6a\x5e\x2e\x4e\x88\x11\x28\xe2\xd5\xd5\x4a\xd5\x4a\x20\x4a\xcf\x37\xb3\x38\x59\x2f\x2c\x31\xa7\x34\x15\x22\x5c\xab\x54\x5b\xcb\xcb\x55\x0b\x08\x00\x00\xff\xff\xc4\xed\x32\xe9\x84\x00\x00\x00")

func values_tmpl() ([]byte, error) {
	return bindata_read(
		_values_tmpl,
		"values.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"bindata.go":  bindata_go,
	"getter.tmpl": getter_tmpl,
	"value.tmpl":  value_tmpl,
	"values.tmpl": values_tmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bindata.go":  &_bintree_t{bindata_go, map[string]*_bintree_t{}},
	"getter.tmpl": &_bintree_t{getter_tmpl, map[string]*_bintree_t{}},
	"value.tmpl":  &_bintree_t{value_tmpl, map[string]*_bintree_t{}},
	"values.tmpl": &_bintree_t{values_tmpl, map[string]*_bintree_t{}},
}}
