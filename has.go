package gotree

import (
	"fmt"
	"reflect"
)

// findHelper recursively searches a node and returns true if any of the node
// satifies the filter else returns false
func hasHelper(node Node, filter FilterFunc) bool {
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
				return true
			}
			if hasHelper(childNode, filter) {
				return true
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
				return true
			}
			if hasHelper(childNode, filter) {
				return true
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
				return true
			}
			if hasHelper(childNode, filter) {
				return true
			}
		}
	case reflect.Interface:
		node.Value = node.Value.Elem()
		return hasHelper(node, filter)
	default:
		node := newNode(node.FullKey, node.Key, node.Value)
		if filter(node) {
			return true
		}
	}

	return false
}

// Has returns true if any node in the tree satifies the filter. It performs a
// depth-first search through the provided data structure and stops at the first
// matching value.
// If no value matches the filter function, it returns false.
//
// Parameters:
//   - tree: The data structure to search (can be a map, slice, array, struct or
//     primitive value)
//   - filter: A function that determines if a value matches the search criteria
//
// Returns:
//   - The first matching value, or error if no match is found or tree is nil.
func Has(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, filter)
}

// HasString searches for the first string value that matches the filter.
// Returns true if any node satifies the filter else returns false.
func HasString(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, FilterString(filter))
}

// HasBool searches for the first bool value that matches the filter. Returns
// true if any node satifies the filter else returns false.
func HasBool(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, FilterBool(filter))
}

// HasInt searches for the first int value that matches the filter. Returns
// true if any node satifies the filter else returns false.
func HasInt(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, FilterInt(filter))
}

// HasInt searches for the first uint value that matches the filter. Returns
// true if any node satifies the filter else returns false.
func HasUInt(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, FilterUint(filter))
}

// HasFloat searches for the first float value that matches the filter. Returns
// true if any node satifies the filter else returns false.
func HasFloat(tree any, filter FilterFunc) bool {
	if tree == nil {
		return false
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return hasHelper(node, FilterFloat(filter))
}
