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
