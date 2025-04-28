package gotree

import (
	"reflect"
	"strings"
	"testing"
)

// EqualSlices checks if two slices have the same elements, regardless of order
func EqualSlices[E comparable](t *testing.T, a, b []E) bool {
	if len(a) != len(b) {
		t.Logf("Unmatched length a = %d, b = %d", len(a), len(b))
		return false
	}
	if len(a) == 0 {
		return true
	}

	counts := make(map[E]int)

	for _, item := range a {
		counts[item]++
	}

	for _, item := range b {
		if counts[item] == 0 {
			t.Log("Missing item", item)
			return false
		}
		counts[item]--
	}

	return true
}

func TestTraverseFunctions(t *testing.T) {
	t.Run("TestTraverse", func(t *testing.T) {
		tests := []struct {
			name   string
			tree   any
			filter func(Node) bool
			want   []any
		}{
			{
				name: "Traverse all strings",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Kind() == reflect.String
				},
				want: []any{
					"hello world",
					"nested value",
					"array_string",
					"found me",
					"Alice",
					"Bob",
				},
			},
			{
				name: "Traverse all integers",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Kind() == reflect.Int ||
						n.Value.Kind() == reflect.Int32 ||
						n.Value.Kind() == reflect.Int64
				},
				want: []any{42, int32(32), int64(64), 100, 123, 999, 30, 25},
			},
			{
				name: "Traverse by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "name"
				},
				want: []any{"Alice", "Bob"},
			},
			{
				name: "Traverse with specific condition",
				tree: testData,
				filter: func(n Node) bool {
					if n.Value.Kind() == reflect.Float64 {
						return n.Value.Float() > 100.0
					}
					return false
				},
				want: []any{123.456, 1250.50, 750.25},
			},
			{
				name: "Traverse empty result",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_key"
				},
				want: []any{},
			},
			{
				name: "Traverse nil tree",
				tree: nil,
				filter: func(n Node) bool {
					return true
				},
				want: []any{},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Traverse(tt.tree, tt.filter)

				// Check length
				if len(got) != len(tt.want) {
					t.Errorf("Traverse() length = %v, want %v", len(got), len(tt.want))
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})

	t.Run("TestTraverseString", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			want      []string
			wantFound bool
		}{
			{
				name:   "Traverse all strings",
				tree:   testData,
				filter: NoneFilter,
				want: []string{
					"hello world",
					"nested value",
					"array_string",
					"found me",
					"Alice",
					"Bob",
				},
				wantFound: true,
			},
			{
				name: "Traverse strings by containing 'e'",
				tree: testData,
				filter: func(n Node) bool {
					return strings.Contains(n.Value.String(), "e")
				},
				want: []string{
					"hello world",
					"nested value",
					"found me",
					"Alice",
				},
				wantFound: true,
			},
			{
				name: "Traverse strings by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "name"
				},
				want:      []string{"Alice", "Bob"},
				wantFound: true,
			},
			{
				name: "Traverse strings empty result",
				tree: testData,
				filter: func(n Node) bool {
					return strings.Contains(n.Value.String(), "zzzzz")
				},
				want:      []string{},
				wantFound: false,
			},
			{
				name:      "Traverse strings nil tree",
				tree:      nil,
				filter:    NoneFilter,
				want:      []string{},
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, gotFound := TraverseString(tt.tree, tt.filter)

				if gotFound != tt.wantFound {
					t.Errorf("TraverseString() found = %v, want %v", gotFound, tt.wantFound)
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})

	t.Run("TestTraverseInt", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			want      []int64
			wantFound bool
		}{
			{
				name:      "Traverse all integers",
				tree:      testData,
				filter:    NoneFilter,
				want:      []int64{42, 32, 64, 100, 123, 999, 30, 25},
				wantFound: true,
			},
			{
				name: "Traverse integers greater than 50",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Int() > 50
				},
				want:      []int64{64, 100, 123, 999},
				wantFound: true,
			},
			{
				name: "Traverse integers by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "age"
				},
				want:      []int64{30, 25},
				wantFound: true,
			},
			{
				name: "Traverse integers empty result",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Int() > 1000
				},
				want:      []int64{},
				wantFound: false,
			},
			{
				name:      "Traverse integers nil tree",
				tree:      nil,
				filter:    NoneFilter,
				want:      []int64{},
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, gotFound := TraverseInt(tt.tree, tt.filter)

				if gotFound != tt.wantFound {
					t.Errorf("TraverseInt() found = %v, want %v", gotFound, tt.wantFound)
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})

	t.Run("TestTraverseFloat", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			want      []float64
			wantFound bool
		}{
			{
				name:      "Traverse all floats",
				tree:      testData,
				filter:    NoneFilter,
				want:      []float64{float64(float32(3.14)), 6.28, 99.99, 45.67, 123.456, 1250.50, 750.25},
				wantFound: true,
			},
			{
				name: "Traverse floats greater than 100",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Float() > 100
				},
				want:      []float64{123.456, 1250.50, 750.25},
				wantFound: true,
			},
			{
				name: "Traverse floats by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "balance"
				},
				want:      []float64{1250.50, 750.25},
				wantFound: true,
			},
			{
				name: "Traverse floats empty result",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Float() > 2000
				},
				want:      []float64{},
				wantFound: false,
			},
			{
				name:      "Traverse floats nil tree",
				tree:      nil,
				filter:    NoneFilter,
				want:      []float64{},
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, gotFound := TraverseFloat(tt.tree, tt.filter)

				if gotFound != tt.wantFound {
					t.Errorf("TraverseFloat() found = %v, want %v", gotFound, tt.wantFound)
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})

	t.Run("TestTraverseBool", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			want      []bool
			wantFound bool
		}{
			{
				name:      "Traverse all booleans",
				tree:      testData,
				filter:    NoneFilter,
				want:      []bool{true, true, false, false, true, false},
				wantFound: true,
			},
			{
				name: "Traverse only true booleans",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Bool() == true
				},
				want:      []bool{true, true, true},
				wantFound: true,
			},
			{
				name: "Traverse booleans by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "active"
				},
				want:      []bool{true, false},
				wantFound: true,
			},
			{
				name: "Traverse booleans empty result",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "non_existent_bool"
				},
				want:      []bool{},
				wantFound: false,
			},
			{
				name:      "Traverse booleans nil tree",
				tree:      nil,
				filter:    NoneFilter,
				want:      []bool{},
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, gotFound := TraverseBool(tt.tree, tt.filter)

				if gotFound != tt.wantFound {
					t.Errorf("TraverseBool() found = %v, want %v", gotFound, tt.wantFound)
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})

	t.Run("TestTraverseUint", func(t *testing.T) {
		tests := []struct {
			name      string
			tree      any
			filter    func(Node) bool
			want      []uint64
			wantFound bool
		}{
			{
				name:      "Traverse all uints",
				tree:      testData,
				filter:    NoneFilter,
				want:      []uint64{10, 64, 200, 88, 888, 1001, 1002},
				wantFound: true,
			},
			{
				name: "Traverse uints greater than 100",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Uint() > 100
				},
				want:      []uint64{200, 888, 1001, 1002},
				wantFound: true,
			},
			{
				name: "Traverse uints by key",
				tree: testData,
				filter: func(n Node) bool {
					return n.Key == "id"
				},
				want:      []uint64{1001, 1002},
				wantFound: true,
			},
			{
				name: "Traverse uints empty result",
				tree: testData,
				filter: func(n Node) bool {
					return n.Value.Uint() > 2000
				},
				want:      []uint64{},
				wantFound: false,
			},
			{
				name:      "Traverse uints nil tree",
				tree:      nil,
				filter:    NoneFilter,
				want:      []uint64{},
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, gotFound := TraverseUint(tt.tree, tt.filter)

				if gotFound != tt.wantFound {
					t.Errorf("TraverseUint() found = %v, want %v", gotFound, tt.wantFound)
					return
				}

				if !EqualSlices(t, tt.want, got) {
					t.Errorf("Traverse() = (%+v), want = (%+v)", got, tt.want)
				}
			})
		}
	})
}
