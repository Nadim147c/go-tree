package main

import (
	"fmt"

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

	var filter gotree.FilterFunc = func(n gotree.Node) bool {
		return n.FullKey == "user.name"
	}
	result, err := gotree.FindString(data, filter)
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // Output: Ephemeral
}
