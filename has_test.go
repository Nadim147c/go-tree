package gotree

import (
	"reflect"
	"strings"
	"testing"
)

func TestHasFunctions(t *testing.T) {
	t.Run("TestHas", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has string by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: true,
			},
			{
				name: "Has deep nested value",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_int"
				},
				want: true,
			},
			{
				name: "Has in array by path",
				tree: testData,
				filter: func(n Node) bool {
					return n.FullKey == "array[4]"
				},
				want: true,
			},
			{
				name: "Has in array nested map",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "array_nested_string"
				},
				want: true,
			},
			{
				name: "Has user by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "name" && n.Value.Kind() == reflect.String {
						return n.Value.String() == "Bob"
					}
					return false
				},
				want: true,
			},
			{
				name: "Has by value condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "age" && n.Value.Kind() == reflect.Int {
						return n.Value.Int() < 30
					}
					return false
				},
				want: true,
			},
			{
				name: "Not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_key"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Has(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("Has() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestHasString", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has existing string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: true,
			},
			{
				name: "Has nested string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_string"
				},
				want: true,
			},
			{
				name: "Has string in array nested map",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "array_nested_string"
				},
				want: true,
			},
			{
				name: "Has string by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "name" {
						return strings.Contains(n.Value.String(), "ice")
					}
					return false
				},
				want: true,
			},
			{
				name: "String not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_string"
				},
				want: false,
			},
			{
				name: "Value exists but not string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int_val"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := HasString(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("HasString() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestHasInt", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has existing int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int_val"
				},
				want: true,
			},
			{
				name: "Has int32",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int32_val"
				},
				want: true,
			},
			{
				name: "Has int64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int64_val"
				},
				want: true,
			},
			{
				name: "Has nested int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_int"
				},
				want: true,
			},
			{
				name: "Has int by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "age" && n.Value.Kind() == reflect.Int {
						return n.Value.Int() > 25
					}
					return false
				},
				want: true,
			},
			{
				name: "Int not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_int"
				},
				want: false,
			},
			{
				name: "Value exists but not int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := HasInt(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("HasInt() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestHasFloat", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has float32",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "float32_val"
				},
				want: true,
			},
			{
				name: "Has float64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "float64_val"
				},
				want: true,
			},
			{
				name: "Has nested float",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_float"
				},
				want: true,
			},
			{
				name: "Has float by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "balance" && n.Value.Kind() == reflect.Float64 {
						return n.Value.Float() > 1000
					}
					return false
				},
				want: true,
			},
			{
				name: "Float not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_float"
				},
				want: false,
			},
			{
				name: "Value exists but not float",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := HasFloat(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("HasFloat() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestHasBool", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has existing bool true",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "bool_val"
				},
				want: true,
			},
			{
				name: "Has nested bool false",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_bool"
				},
				want: true,
			},
			{
				name: "Has bool by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "active" && n.Value.Kind() == reflect.Bool {
						return !n.Value.Bool() // Find inactive user
					}
					return false
				},
				want: true,
			},
			{
				name: "Bool not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_bool"
				},
				want: false,
			},
			{
				name: "Value exists but not bool",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := HasBool(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("HasBool() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestHasUInt", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   bool
		}{
			{
				name: "Has existing uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "uint_val"
				},
				want: true,
			},
			{
				name: "Has uint64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "uint64_val"
				},
				want: true,
			},
			{
				name: "Has nested uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_uint"
				},
				want: true,
			},
			{
				name: "Has uint by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "id" && n.Value.Kind() == reflect.Uint64 {
						return n.Value.Uint() > 1001
					}
					return false
				},
				want: true,
			},
			{
				name: "Uint not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_uint"
				},
				want: false,
			},
			{
				name: "Value exists but not uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: false,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := HasUInt(tt.tree, tt.filter)
				if got != tt.want {
					t.Errorf("HasUInt() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
