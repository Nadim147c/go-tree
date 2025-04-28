package gotree

import (
	"errors"
	"reflect"
)

var (
	ErrNilTree  = errors.New("tree is nil")
	ErrNotFound = errors.New("No item found")
)

// FilterFunc defines a function type that takes a Node and returns a boolean
// value indicating whether the node satisfies certain conditions.
type FilterFunc func(Node) bool

func NoneFilter(_ Node) bool {
	return true
}

// FilterString returns a FilterFunc that checks if a node's value is a string
// type and satisfies the provided filter condition.
func FilterString(filter FilterFunc) FilterFunc {
	return func(n Node) bool {
		return n.Value.Kind() == reflect.String && filter(n)
	}
}

// FilterBool returns a FilterFunc that checks if a node's value is a boolean
// type and satisfies the provided filter condition.
func FilterBool(filter FilterFunc) FilterFunc {
	return func(n Node) bool {
		return n.Value.Kind() == reflect.Bool && filter(n)
	}
}

// FilterInt returns a FilterFunc that checks if a node's value is any integer
// type (int, int8, int16, int32, or int64) and satisfies the provided filter
// condition.
func FilterInt(filter FilterFunc) FilterFunc {
	return func(n Node) bool {
		switch n.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return filter(n)
		default:
			return false
		}
	}
}

// FilterUint returns a FilterFunc that checks if a node's value is any unsigned
// integer type (uint, uint8, uint16, uint32, or uint64) and satisfies the
// provided filter condition.
func FilterUint(filter FilterFunc) FilterFunc {
	return func(n Node) bool {
		switch n.Value.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return filter(n)
		default:
			return false
		}
	}
}

// FilterFloat returns a FilterFunc that checks if a node's value is a floating
// point type (float32 or float64) and satisfies the provided filter condition.
func FilterFloat(filter FilterFunc) FilterFunc {
	return func(n Node) bool {
		switch n.Value.Kind() {
		case reflect.Float32, reflect.Float64:
			return filter(n)
		default:
			return false
		}
	}
}
