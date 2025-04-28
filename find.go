package gotree

import (
	"fmt"
	"reflect"
)

// findHelper recursively searches a node and returns the first value that
// matches the filter function. It traverses maps, slices, arrays, structs, and
// interfaces through type reflection. When a matching node is found (filter
// returns true), it immediately returns that value and stops traversal.
func findHelper(node Node, filter func(Node) bool) (Node, bool) {
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
				return childNode, true
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
				return childNode, true
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
				return childNode, true
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
			return node, true
		}
	}
	return Node{}, false
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
//   - The first matching value, or error if no match is found or tree is nil.
func Find(tree any, filter func(Node) bool) (any, error) {
	if tree == nil {
		return nil, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	if v, exists := findHelper(node, filter); exists {
		return v.Interface, nil
	}

	return nil, ErrNotFound
}

// FindString searches for the first string value that matches the filter.
// Returns the string if found, otherwise returns an error.
func FindString(tree any, filter func(Node) bool) (string, error) {
	if tree == nil {
		return "", ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	val, ok := findHelper(node, FilterString(filter))
	if !ok || val.Interface == nil {
		return "", ErrNotFound
	}

	return val.Value.String(), nil
}

// FindBool searches for the first bool value that matches the filter. Returns
// the bool if found, otherwise returns an error.
func FindBool(tree any, filter func(Node) bool) (bool, error) {
	if tree == nil {
		return false, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	val, ok := findHelper(node, FilterBool(filter))
	if !ok || val.Interface == nil {
		return false, ErrNotFound
	}
	return val.Value.Bool(), nil
}

// FindInt searches for the first int value that matches the filter. Returns the
// int64 if found, otherwise returns an error.
func FindInt(tree any, filter func(Node) bool) (int64, error) {
	if tree == nil {
		return 0, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	val, ok := findHelper(node, FilterInt(filter))
	if !ok || val.Interface == nil {
		return 0, ErrNotFound
	}
	return val.Value.Int(), nil
}

// FindUint searches for the first uint value that matches the filter. Returns
// the uint64 if found, otherwise returns an error.
func FindUint(tree any, filter func(Node) bool) (uint64, error) {
	if tree == nil {
		return 0, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	val, ok := findHelper(node, FilterUint(filter))
	if !ok || val.Interface == nil {
		return 0, ErrNotFound
	}
	return val.Value.Uint(), nil
}

// FindFloat searches for the first float value that matches the filter. Returns
// the float64 if found, otherwise returns an error.
func FindFloat(tree any, filter func(Node) bool) (float64, error) {
	if tree == nil {
		return 0, ErrNilTree
	}

	node := newNode("", "", reflect.ValueOf(tree))
	val, ok := findHelper(node, FilterFloat(filter))
	if !ok || val.Interface == nil {
		return 0, ErrNotFound
	}
	return val.Value.Float(), nil
}
