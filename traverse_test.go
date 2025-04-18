package gotree_test

import (
	"reflect"
	"strings"
	"testing"

	gotree "github.com/Nadim147c/go-tree"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name   string
		tree   any
		filter func(gotree.Node) bool
		want   []any
	}{
		{
			name: "Simple map with strings",
			tree: map[string]string{
				"name":  "John",
				"email": "john@example.com",
			},
			filter: func(n gotree.Node) bool {
				return n.Key != "" // Include all non-root nodes
			},
			want: []any{"John", "john@example.com"},
		},
		{
			name: "Nested map",
			tree: map[string]any{
				"user": map[string]any{
					"name": "Alice",
					"address": map[string]string{
						"street": "123 Main St",
						"city":   "Anytown",
					},
				},
			},
			filter: func(n gotree.Node) bool {
				return strings.HasPrefix(n.FullKey, "user.address.")
			},
			want: []any{"123 Main St", "Anytown"},
		},
		{
			name: "Filter by value type",
			tree: map[string]any{
				"id":    42,
				"name":  "Product",
				"price": 19.99,
				"tags":  []string{"sale", "featured"},
			},
			filter: func(n gotree.Node) bool {
				return n.Value.Kind() == reflect.Float64
			},
			want: []any{19.99},
		},
		{
			name: "Array of primitives",
			tree: []int{1, 2, 3, 4, 5},
			filter: func(n gotree.Node) bool {
				if n.Value.Kind() == reflect.Int {
					return n.Value.Int()%2 == 0 // Only even numbers
				}
				return false
			},
			want: []any{2, 4},
		},
		{
			name: "Array of objects",
			tree: []map[string]any{
				{"id": 1, "active": true},
				{"id": 2, "active": false},
				{"id": 3, "active": true},
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "active" && n.Value.Kind() == reflect.Bool && n.Value.Bool()
			},
			want: []any{true, true},
		},
		{
			name: "Complex nested structure",
			tree: map[string]any{
				"company": map[string]any{
					"name": "Acme Corp",
					"departments": []map[string]any{
						{
							"name":     "Engineering",
							"staff":    15,
							"location": "Floor 3",
						},
						{
							"name":     "Marketing",
							"staff":    8,
							"location": "Floor 2",
						},
					},
					"founded": 1985,
				},
			},
			filter: func(n gotree.Node) bool {
				return strings.Contains(n.FullKey, ".departments[") && n.Value.Kind() == reflect.String && n.Key == "name"
			},
			want: []any{"Engineering", "Marketing"},
		},
		{
			name: "Empty tree",
			tree: map[string]string{},
			filter: func(n gotree.Node) bool {
				return true
			},
			want: []any{},
		},
		{
			name: "Nil tree",
			tree: nil,
			filter: func(n gotree.Node) bool {
				return true
			},
			want: []any{},
		},
		{
			name: "Filter includes nothing",
			tree: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			filter: func(n gotree.Node) bool {
				return false // Filter out everything
			},
			want: []any{},
		},
		{
			name: "Primitive value as root",
			tree: "just a string",
			filter: func(n gotree.Node) bool {
				return n.Value.Kind() == reflect.String
			},
			want: []any{"just a string"},
		},
		{
			name: "Deep nesting",
			tree: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{
						"level3": map[string]any{
							"level4": map[string]any{
								"target": "found me!",
							},
						},
					},
				},
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "target"
			},
			want: []any{"found me!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gotree.Traverse(tt.tree, tt.filter)

			// Compare slices without caring about order
			if !compareSlicesIgnoreOrder(got, tt.want) {
				t.Errorf("Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// compareSlicesIgnoreOrder compares two slices for equality without considering order
func compareSlicesIgnoreOrder(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each element
	countA := make(map[any]int)
	countB := make(map[any]int)

	for _, v := range a {
		countA[v]++
	}

	for _, v := range b {
		countB[v]++
	}

	// Check that counts match for all elements
	for k, v := range countA {
		if countB[k] != v {
			return false
		}
	}

	return true
}
