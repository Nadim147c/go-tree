package main

import (
	"fmt"
	"reflect"

	"github.com/Nadim147c/go-tree"
)

func main() {
	data := map[string]any{
		"users": []any{
			map[string]any{"name": "Alice"},
			map[string]any{"name": "Bob"},
		},
	}

	results := gotree.Traverse(data, func(n gotree.Node) bool {
		return n.Key == "name" && n.Value.Kind() == reflect.String
	})

	for _, r := range results {
		fmt.Println(r)
	}
	// Output:
	// Alice
	// Bob
}
