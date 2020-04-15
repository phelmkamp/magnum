# magnum
Go enum generator

# Installation

`go get github.com/phelmkamp/magnum`

# Usage

1. Define a simple struct type
    1. Enumerate values in the struct tag of the `name` field

        Format is `enum:"value[,value2]"`:
        ```go
        type Color struct {
            name string `enum:"red,orange,yellow,green,blue,indigo,violet"`
        }
        ```
        
    2. (Optional) Enumerate values of other fields with additional struct tags (see [house.House](testdata/house/house.go))

2. Run command

	```bash
	magnum --path=$SRCDIR
	```

	Better yet, add the following comment to a file at the root of your source tree (e.g. main.go)
	and run `go generate` as part of your build process

	```go
	//go:generate magnum
	```

3. That's it!

	An *_enum.go file is generated for each *.go file that has enum tags. This file contains accessors for the enumerated values plus methods to marshal/unmarshal a value to/from text (and JSON).
