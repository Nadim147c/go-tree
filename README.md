# go-tree

[![Go Reference](https://pkg.go.dev/badge/github.com/Nadim147c/go-tree.svg)](https://pkg.go.dev/github.com/Nadim147c/go-tree)

`go-tree` is a simple Go package for traversing and searching deeply nested data structures like
maps, slices, arrays, and structs. It provides two powerful functions`Find` and `Traverse`that make
it easy to extract specific values from complex JSON-like trees using custom filter logic.

## Installation

```bash
go get github.com/Nadim147c/go-tree
```

## Features

- Depth-first traversal of nested structures
- Flexible filtering with custom logic
- Supports maps, slices, arrays, structs, and primitive types

## Usage

Check the [example](./example/) directory for example.

#### Find(tree any, filter func(Node) bool) any
Returns the first value in the data structure that matches the filter.

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/Nadim147c/go-tree"
)

func main() {
	data := map[string]any{
		"user": map[string]any{
			"name": "Ephemeral",
			"age":  30,
		},
		"active": true,
	}

	result := gotree.Find(data, func(n gotree.Node) bool {
		return n.Key == "name" && n.Value.Kind() == reflect.String
	})

	fmt.Println(result) // Output: Ephemeral
}
```

> Note: Make sure to check type of the value in filter function by `n.Value.Kind() ==
> reflect.YourType` for primitive types.


## Node Structure

The `Node` struct provides rich context during traversal:

```go
type Node struct {
	FullKey   string        // Full path from the root, e.g., "users[0].name"
	Key       string        // Immediate key or index
	Value     reflect.Value // Raw reflect.Value of the node
	Interface any           // Value as interface{}
}
```

## License

This package is licensed under [GNU-LGPL-3.0](./LICENSE).
