package gotree

import (
	"fmt"
	"reflect"
)

// findHelper recursively searches a node and returns the first value that
// matches the filter function. It traverses maps, slices, arrays, structs, and
// interfaces through type reflection. When a matching node is found (filter
// returns true), it immediately returns that value and stops traversal.
func findHelper(node Node, filter func(Node) bool) (any, bool) {
	switch node.Value.Kind() {
	case reflect.Map:
		for _, k := range node.Value.MapKeys() {
			strKey := fmt.Sprint(k.Interface())
			newFullKey := strKey
			if node.FullKey != "" {
				newFullKey = node.FullKey + "." + strKey
			}

			v := node.Value.MapIndex(k)
			childNode := newNode(newFullKey, strKey, v)
			if filter(childNode) {
				return v.Interface(), true
			}
			if result, found := findHelper(childNode, filter); found {
				return result, true
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < node.Value.Len(); i++ {
			strKey := fmt.Sprintf("[%d]", i)
			newFullKey := strKey
			if node.FullKey != "" {
				newFullKey = node.FullKey + strKey
			}

			v := node.Value.Index(i)
			childNode := newNode(newFullKey, strKey, v)
			if filter(childNode) {
				return v.Interface(), true
			}
			if result, found := findHelper(childNode, filter); found {
				return result, true
			}
		}
	case reflect.Struct:
		reflectType := node.Value.Type()
		for i := 0; i < node.Value.NumField(); i++ {
			field := reflectType.Field(i)
			if !field.IsExported() {
				continue
			}

			newFullKey := field.Name
			if node.FullKey != "" {
				newFullKey = node.FullKey + "." + field.Name
			}

			v := node.Value.Field(i)
			childNode := newNode(newFullKey, field.Name, v)
			if filter(childNode) {
				return v.Interface(), true
			}
			if result, found := findHelper(childNode, filter); found {
				return result, true
			}
		}
	case reflect.Interface:
		node.Value = node.Value.Elem()
		return findHelper(node, filter)
	default:
		node := newNode(node.FullKey, node.Key, node.Value)
		if filter(node) {
			return node.Interface, true
		}
	}
	return nil, false
}

// Find returns the first value that matches the given filter function. It
// performs a depth-first search through the provided data structure and stops
// at the first matching value.
//
// If no value matches the filter function, it returns nil.
//
// Parameters:
//   - tree: The data structure to search (can be a map, slice, array, struct or
//     primitive value)
//   - filter: A function that determines if a value matches the search criteria
//
// Returns:
//   - The first matching value, or nil if no match is found
func Find(tree any, filter func(Node) bool) any {
	if tree == nil {
		return nil
	}

	node := newNode("", "", reflect.ValueOf(tree))
	if v, exists := findHelper(node, filter); exists {
		return v
	}
	return nil
}
