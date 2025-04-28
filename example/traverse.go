package main

import (
	"fmt"

	"github.com/Nadim147c/go-tree"
)

func main() {
	data := map[string]any{
		"users": []any{
			map[string]any{"name": "Alice"},
			map[string]any{"name": "Bob"},
		},
	}

	var filter gotree.FilterFunc = func(n gotree.Node) bool {
		return n.Key == "name"
	}

	results, err := gotree.TraverseString(data, filter)
	if err != nil {
		panic(err)
	}

	for _, r := range results {
		fmt.Println(r)
	}
	// Output:
	// Alice
	// Bob
}
