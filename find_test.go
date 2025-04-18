package gotree_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Nadim147c/go-tree"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name   string
		tree   any
		filter func(gotree.Node) bool
		want   any
	}{
		{
			name: "Find in simple map",
			tree: map[string]string{
				"name":  "John",
				"email": "john@example.com",
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "email"
			},
			want: "john@example.com",
		},
		{
			name: "Find in nested map",
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
				return n.Key == "street"
			},
			want: "123 Main St",
		},
		{
			name: "Find by full path",
			tree: map[string]any{
				"users": []map[string]any{
					{
						"id":   1,
						"name": "User 1",
					},
					{
						"id":   2,
						"name": "User 2",
					},
				},
			},
			filter: func(n gotree.Node) bool {
				// Second user's name
				return n.FullKey == "users[1].name" && n.Value.Kind() == reflect.String
			},
			want: "User 2",
		},
		{
			name: "Find by value type and condition",
			tree: map[string]any{
				"items": []map[string]any{
					{"id": 1, "price": 10.5},
					{"id": 2, "price": 20.75},
					{"id": 3, "price": 30.0},
				},
			},
			filter: func(n gotree.Node) bool {
				if n.Key == "price" && n.Value.Kind() == reflect.Float64 {
					return n.Value.Float() > 20.0
				}
				return false
			},
			want: 20.75, // First price over 20.0
		},
		{
			name: "Find in array of primitives",
			tree: []int{1, 2, 3, 4, 5},
			filter: func(n gotree.Node) bool {
				if n.Value.Kind() == reflect.Int {
					return n.Value.Int() > 3 // First number > 3
				}
				return false
			},
			want: 4,
		},
		{
			name: "Complex nested search",
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
				},
			},
			filter: func(n gotree.Node) bool {
				if n.Key == "staff" && n.Value.Kind() == reflect.Int {
					return n.Value.Int() < 10 // Department with less than 10 staff
				}
				return false
			},
			want: 8, // Marketing department has 8 staff
		},
		{
			name: "Not found",
			tree: map[string]any{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "z" // Key doesn't exist
			},
			want: nil,
		},
		{
			name: "Empty tree",
			tree: map[string]string{},
			filter: func(n gotree.Node) bool {
				return true
			},
			want: nil,
		},
		{
			name: "Nil tree",
			tree: nil,
			filter: func(n gotree.Node) bool {
				return true
			},
			want: nil,
		},
		{
			name: "Find in primitive root",
			tree: "test string",
			filter: func(n gotree.Node) bool {
				return n.Value.Kind() == reflect.String
			},
			want: "test string",
		},
		{
			name: "Deep nesting find",
			tree: map[string]any{
				"level1": map[string]any{
					"level2": map[string]any{
						"level3": map[string]any{
							"level4": map[string]any{
								"needle": "found in haystack",
							},
						},
					},
				},
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "needle"
			},
			want: "found in haystack",
		},
		{
			name: "Find first in multiple matches",
			tree: map[string]any{
				"users": []map[string]any{
					{"status": "active", "name": "Alice"},
					{"status": "active", "name": "Bob"},
					{"status": "inactive", "name": "Charlie"},
				},
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "status" && n.Interface == "active"
			},
			want: "active", // Should return the first active status (Alice's)
		},
		{
			name: "Case sensitive key search",
			tree: map[string]any{
				"Name": "Upper",
				"name": "Lower",
			},
			filter: func(n gotree.Node) bool {
				return n.Key == "name" // Should match lowercase only
			},
			want: "Lower",
		},
		{
			name: "Find using string contains",
			tree: map[string]any{
				"users": []map[string]any{
					{"email": "alice@example.com"},
					{"email": "bob@test.org"},
					{"email": "charlie@example.com"},
				},
			},
			filter: func(n gotree.Node) bool {
				if n.Key == "email" && n.Value.Kind() == reflect.String {
					return strings.Contains(n.Value.String(), "test.org")
				}
				return false
			},
			want: "bob@test.org",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gotree.Find(tt.tree, tt.filter)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
