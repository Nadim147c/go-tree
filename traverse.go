package gotree

import (
	"fmt"
	"reflect"
)

// traverseHelper recursively traverses a node and collects values based on the
// filter function. It handles maps, slices, arrays, structs, and interfaces
// through type reflection. For each node, it either collects the value (if
// filter returns true) or continues traversing deeper.
func traverseHelper(node Node, filter FilterFunc) []Node {
	results := make([]Node, 0)

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
				results = append(results, childNode)
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
				results = append(results, childNode)
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
				results = append(results, childNode)
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
			results = append(results, node)
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
//   - The a slice matching value, or error if no match is found or tree is nil.
func Traverse(tree any, filter FilterFunc) ([]any, error) {
	if tree == nil {
		return []any{}, ErrNilTree
	}
	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, filter)

	if len(nodes) == 0 {
		return []any{}, ErrNotFound
	}

	values := make([]any, len(nodes))
	for i, v := range nodes {
		values[i] = v.Interface
	}
	return values, nil
}

// TraverseString searches for all string values in the tree that match the
// filter. Returns a slice of matching string values and an error if none found.
func TraverseString(tree any, filter FilterFunc) ([]string, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, FilterString(filter))

	if len(nodes) == 0 {
		return nil, ErrNotFound
	}

	values := make([]string, len(nodes))
	for i, v := range nodes {
		values[i] = v.Value.String()
	}

	return values, nil
}

// TraverseBool searches for all boolean values in the tree that match the
// filter. Returns a slice of matching boolean values and an error if none
// found.
func TraverseBool(tree any, filter FilterFunc) ([]bool, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, FilterBool(filter))

	if len(nodes) == 0 {
		return nil, ErrNotFound
	}

	values := make([]bool, len(nodes))
	for i, v := range nodes {
		values[i] = v.Value.Bool()
	}

	return values, nil
}

// TraverseInt searches for all integer values in the tree that match the
// filter. Returns a slice of matching integer values and an error if none
// found.
func TraverseInt(tree any, filter FilterFunc) ([]int64, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, FilterInt(filter))

	if len(nodes) == 0 {
		return nil, ErrNotFound
	}

	values := make([]int64, len(nodes))
	for i, v := range nodes {
		values[i] = v.Value.Int()
	}

	return values, nil
}

// TraverseUint searches for all unsigned integer values in the tree that match
// the filter. Returns a slice of matching unsigned integer values and an error
// if none found.
func TraverseUint(tree any, filter FilterFunc) ([]uint64, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, FilterUint(filter))

	if len(nodes) == 0 {
		return nil, ErrNotFound
	}

	values := make([]uint64, len(nodes))
	for i, v := range nodes {
		values[i] = v.Value.Uint()
	}

	return values, nil
}

// TraverseFloat searches for all floating point values in the tree that match
// the filter. Returns a slice of matching float values and an error if none
// found.
func TraverseFloat(tree any, filter FilterFunc) ([]float64, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	nodes := traverseHelper(node, FilterFloat(filter))

	if len(nodes) == 0 {
		return nil, ErrNotFound
	}

	values := make([]float64, len(nodes))
	for i, v := range nodes {
		values[i] = v.Value.Float()
	}

	return values, nil
}
