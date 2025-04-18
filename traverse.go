package gotree

import (
	"fmt"
	"reflect"
)

// traverseHelper recursively traverses a node and collects values based on the
// filter function. It handles maps, slices, arrays, structs, and interfaces
// through type reflection. For each node, it either collects the value (if
// filter returns true) or continues traversing deeper.
func traverseHelper(node Node, filter func(Node) bool) []any {
	results := make([]any, 0)

	switch node.Value.Kind() {
	case reflect.Map:
		// Iterate over all map keys
		for _, k := range node.Value.MapKeys() {
			strKey := fmt.Sprint(k.Interface())
			newFullKey := strKey
			if node.FullKey != "" {
				newFullKey = node.FullKey + "." + strKey
			}

			v := node.Value.MapIndex(k)
			childNode := newNode(newFullKey, strKey, v)
			if filter(childNode) {
				results = append(results, v.Interface())
			} else {
				results = append(results, traverseHelper(childNode, filter)...)
			}
		}
	case reflect.Slice, reflect.Array:
		// Iterate over all slice index
		for i := 0; i < node.Value.Len(); i++ {
			strKey := fmt.Sprintf("[%d]", i) // Array key format
			newFullKey := strKey
			if node.FullKey != "" {
				newFullKey = node.FullKey + strKey
			}

			v := node.Value.Index(i)
			childNode := newNode(newFullKey, strKey, v)
			if filter(childNode) {
				results = append(results, v.Interface())
			} else {
				results = append(results, traverseHelper(childNode, filter)...)
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
				results = append(results, v.Interface())
			} else {
				results = append(results, traverseHelper(childNode, filter)...)
			}
		}
	case reflect.Interface:
		node.Value = node.Value.Elem()
		results = append(results, traverseHelper(node, filter)...)
	default:
		node := newNode(node.FullKey, node.Key, node.Value)
		if filter(node) {
			results = append(results, node.Interface)
		}
	}
	return results
}

// Traverse traverses a nested JSON tree and returns all values for which the
// filter function returns true. It initializes the traversal with the root
// node of the data structure and then calls traverseHelper to perform the
// actual recursive traversal.
//
// The filter function receives each node during traversal and determines
// whether to include the node's value in the results.
//
// Parameters:
//   - tree: The data structure to traverse (can be a map, slice, array, struct
//     or primitive value)
//   - filter: A function that determines which values to include in the results
//
// Returns:
//   - A slice containing all values that passed the filter function
func Traverse(tree any, filter func(Node) bool) []any {
	if tree == nil {
		return make([]any, 0)
	}
	node := newNode("", "", reflect.ValueOf(tree))
	return traverseHelper(node, filter)
}
