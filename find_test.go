package gotree

import (
	"reflect"
	"strings"
	"testing"
)

// Sample test data structure
var testData = map[string]any{
	"string_val":  "hello world",
	"int_val":     42,
	"int32_val":   int32(32),
	"int64_val":   int64(64),
	"uint_val":    uint(10),
	"uint64_val":  uint64(64),
	"float32_val": float32(3.14),
	"float64_val": 6.28,
	"bool_val":    true,
	"nested": map[string]any{
		"deep_string": "nested value",
		"deep_int":    100,
		"deep_float":  99.99,
		"deep_bool":   false,
		"deep_uint":   uint64(200),
	},
	"array": []any{
		"array_string",
		123,
		45.67,
		false,
		uint(88),
		map[string]any{
			"array_nested_string": "found me",
			"array_nested_int":    999,
			"array_nested_float":  123.456,
			"array_nested_bool":   true,
			"array_nested_uint":   uint64(888),
		},
	},
	"users": []map[string]any{
		{
			"name":    "Alice",
			"age":     30,
			"balance": 1250.50,
			"active":  true,
			"id":      uint64(1001),
		},
		{
			"name":    "Bob",
			"age":     25,
			"balance": 750.25,
			"active":  false,
			"id":      uint64(1002),
		},
	},
}

func TestFindFunctions(t *testing.T) {
	t.Run("TestFind", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   any
		}{
			{
				name: "Find string by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				want: "hello world",
			},
			{
				name: "Find deep nested value",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_int"
				},
				want: 100,
			},
			{
				name: "Find in array by path",
				tree: testData,
				filter: func(n Node) bool {
					return n.FullKey == "array[4]"
				},
				want: uint(88),
			},
			{
				name: "Find in array nested map",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "array_nested_string"
				},
				want: "found me",
			},
			{
				name: "Find user by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "name" && n.Value.Kind() == reflect.String {
						return n.Value.String() == "Bob"
					}
					return false
				},
				want: "Bob",
			},
			{
				name: "Find by value condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "age" && n.Value.Kind() == reflect.Int {
						return n.Value.Int() < 30
					}
					return false
				},
				want: 25,
			},
			{
				name: "Not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_key"
				},
				want: nil,
			},
			{
				name: "Nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Find(tt.tree, tt.filter)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Find() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("TestFindString", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			wantValue string
			wantFound bool
		}{
			{
				name: "Find existing string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				wantValue: "hello world",
				wantFound: true,
			},
			{
				name: "Find nested string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_string"
				},
				wantValue: "nested value",
				wantFound: true,
			},
			{
				name: "Find string in array nested map",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "array_nested_string"
				},
				wantValue: "found me",
				wantFound: true,
			},
			{
				name: "Find string by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "name" && n.Value.Kind() == reflect.String {
						return strings.Contains(n.Value.String(), "ice")
					}
					return false
				},
				wantValue: "Alice",
				wantFound: true,
			},
			{
				name: "String not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_string"
				},
				wantValue: "",
				wantFound: false,
			},
			{
				name: "Value exists but not string",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int_val"
				},
				wantValue: "",
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue, gotFound := FindString(tt.tree, tt.filter)
				if gotValue != tt.wantValue || gotFound != tt.wantFound {
					t.Errorf("FindString() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
				}
			})
		}
	})

	t.Run("TestFindInt", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			wantValue int64
			wantFound bool
		}{
			{
				name: "Find existing int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int_val"
				},
				wantValue: 42,
				wantFound: true,
			},
			{
				name: "Find int32",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int32_val"
				},
				wantValue: 32,
				wantFound: true,
			},
			{
				name: "Find int64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "int64_val"
				},
				wantValue: 64,
				wantFound: true,
			},
			{
				name: "Find nested int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_int"
				},
				wantValue: 100,
				wantFound: true,
			},
			{
				name: "Find int by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "age" && n.Value.Kind() == reflect.Int {
						return n.Value.Int() > 25
					}
					return false
				},
				wantValue: 30,
				wantFound: true,
			},
			{
				name: "Int not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_int"
				},
				wantValue: 0,
				wantFound: false,
			},
			{
				name: "Value exists but not int",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				wantValue: 0,
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue, gotFound := FindInt(tt.tree, tt.filter)
				if gotValue != tt.wantValue || gotFound != tt.wantFound {
					t.Errorf("FindInt() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
				}
			})
		}
	})

	t.Run("TestFindFloat", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			wantValue float64
			wantFound bool
		}{
			{
				name: "Find float32",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "float32_val"
				},
				wantValue: float64(float32(3.14)),
				wantFound: true,
			},
			{
				name: "Find float64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "float64_val"
				},
				wantValue: 6.28,
				wantFound: true,
			},
			{
				name: "Find nested float",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_float"
				},
				wantValue: 99.99,
				wantFound: true,
			},
			{
				name: "Find float by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "balance" && n.Value.Kind() == reflect.Float64 {
						return n.Value.Float() > 1000
					}
					return false
				},
				wantValue: 1250.50,
				wantFound: true,
			},
			{
				name: "Float not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_float"
				},
				wantValue: 0,
				wantFound: false,
			},
			{
				name: "Value exists but not float",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				wantValue: 0,
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue, gotFound := FindFloat(tt.tree, tt.filter)
				if gotValue != tt.wantValue || gotFound != tt.wantFound {
					t.Errorf("FindFloat() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
				}
			})
		}
	})

	t.Run("TestFindBool", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			wantValue bool
			wantFound bool
		}{
			{
				name: "Find existing bool true",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "bool_val"
				},
				wantValue: true,
				wantFound: true,
			},
			{
				name: "Find nested bool false",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_bool"
				},
				wantValue: false,
				wantFound: true,
			},
			{
				name: "Find bool by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "active" && n.Value.Kind() == reflect.Bool {
						return !n.Value.Bool() // Find inactive user
					}
					return false
				},
				wantValue: false,
				wantFound: true,
			},
			{
				name: "Bool not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_bool"
				},
				wantValue: false,
				wantFound: false,
			},
			{
				name: "Value exists but not bool",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				wantValue: false,
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue, gotFound := FindBool(tt.tree, tt.filter)
				if gotValue != tt.wantValue || gotFound != tt.wantFound {
					t.Errorf("FindBool() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
				}
			})
		}
	})

	t.Run("TestFindUint", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			wantValue uint64
			wantFound bool
		}{
			{
				name: "Find existing uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "uint_val"
				},
				wantValue: 10,
				wantFound: true,
			},
			{
				name: "Find uint64",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "uint64_val"
				},
				wantValue: 64,
				wantFound: true,
			},
			{
				name: "Find nested uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "deep_uint"
				},
				wantValue: 200,
				wantFound: true,
			},
			{
				name: "Find uint by condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Key == "id" && n.Value.Kind() == reflect.Uint64 {
						return n.Value.Uint() > 1001
					}
					return false
				},
				wantValue: 1002,
				wantFound: true,
			},
			{
				name: "Uint not found",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_uint"
				},
				wantValue: 0,
				wantFound: false,
			},
			{
				name: "Value exists but not uint",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "string_val"
				},
				wantValue: 0,
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue, gotFound := FindUint(tt.tree, tt.filter)
				if gotValue != tt.wantValue || gotFound != tt.wantFound {
					t.Errorf("FindUint() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
				}
			})
		}
	})
}
