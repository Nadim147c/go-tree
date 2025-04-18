package gotree

import "reflect"

// Node represents a node in a data structure being traversed. It contains
// information about the node's location in the structure (FullKey and Key),
// its value as a reflect.Value, and its value as an interface{}.
type Node struct {
	// FullKey is the complete path to this node from the root
	// (e.g., "user.address.street")
	FullKey string

	// Key is the immediate key or field name of this node
	// (e.g., "street" in the example above)
	Key string

	// Value is the reflect.Value representation of this node's value
	Value reflect.Value

	// Interface is the node's value as an interface{}
	Interface any
}

// newNode creates a new Node with the given full key path, immediate key, and
// reflect.Value. It automatically extracts the interface{} value from the
// reflect.Value.
func newNode(fullKey, key string, value reflect.Value) Node {
	return Node{
		FullKey:   fullKey,
		Key:       key,
		Value:     value,
		Interface: value.Interface(),
	}
}
