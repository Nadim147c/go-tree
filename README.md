# go-tree

[![Go Reference](https://pkg.go.dev/badge/github.com/Nadim147c/go-tree.svg)](https://pkg.go.dev/github.com/Nadim147c/go-tree)

`go-tree` is a simple Go package for **searching** and **traversing** deeply nested data structures such as maps, slices, arrays, and structs.
It offers flexible utilities for extracting specific values or entire subtrees using custom filter logic.

## Installation

```bash
go get github.com/Nadim147c/go-tree
```

## Features

- Depth-first traversal of nested structures
- Customizable filtering
- Supports maps, slices, arrays, structs, and primitive types
- Clear and predictable error handling

## Usage

See the [example](./example/) directory for practical examples.

## How it Works

`go-tree` provides two main capabilities:

### Find

- **Use `Find<Type>`** (`FindString`, `FindBool`, `FindInt`, `FindUint`, `FindFloat`) when you want to **find a primitive value**.
  These functions automatically **checks the type** before running filter function.
- **Use `Find`** (generic) only when you want to **find a branch** (a nested map, slice, struct, etc.) for **further processing** with `Find<Type>` or `Traverse<Type>`.

### Traverse

- **Use `Traverse<Type>`** (`TraverseString`, `TraverseBool`, `TraverseInt`, `TraverseUint`, `TraverseFloat`) to **collect multiple primitive values** from the structure, with **type checking** before running filter function.
- **Use `Traverse`** (generic) when you want to **collect branches** (complex nested structures) for further analysis.

# Summary

- For **primitive values**, always prefer `Find<Type>` and `Traverse<Type>`.
- Use **generic `Find` and `Traverse`** when working with **nested objects**.
- Full API documentation is available through GoDoc or your IDE.
- See [example](./example/) for practical usage.

## License

This package is licensed under [GNU-LGPL-3.0](./LICENSE).
