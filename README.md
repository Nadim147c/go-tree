# go-tree

[![Go Reference](https://pkg.go.dev/badge/github.com/Nadim147c/go-tree.svg)](https://pkg.go.dev/github.com/Nadim147c/go-tree)

`go-tree` is a simple Go package for **searching** and **traversing** deeply nested data structures
such as maps, slices, arrays, and structs. It offers flexible utilities for extracting specific
values or entire subtrees using custom filter logic.

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

`go-tree` provides following capabilities:

### Find

- `Find<Type>` functions (`FindString`, `FindBool`, `FindInt`, `FindUint`, `FindFloat`) scan the
  tree and return the **first primitive value** of the specified type that satisfies the filter.
- Type checks are performed **before** the filter is applied, ensuring only correctly-typed values
  are evaluated.
- The generic `Find` function returns the first matching **branch** (e.g., `map`, `slice`, `struct`)
  without performing type filtering — useful for targeting nested structures for deeper inspection.

### Traverse

- `Traverse<Type>` functions collect **all primitive values** of the specified type that match the
  filter. These include `TraverseString`, `TraverseBool`, `TraverseInt`, `TraverseUint`, and
  `TraverseFloat`.
- Type enforcement is done before applying the filter, ensuring that only relevant values are
  processed.
- The generic `Traverse` collects **all matching branches**, useful for aggregating nested
  collections of interest, regardless of their concrete type.

### Has

- `Has<Type>` functions (`HasString`, `HasBool`, `HasInt`, `HasUInt`, `HasFloat`) traverse the
  structure and return `true` if **any value of the specified type** satisfies the filter function.
- These functions **automatically perform type checking** before applying the filter, which
  simplifies filter logic but also restricts it to only matching values of the expected type.
- The generic `Has` does **not enforce type constraints**, making it useful when the filter logic
  needs to inspect or match across **multiple types or complex conditions**.

# Summary

- For **primitive values**, always prefer `Find<Type>` and `Traverse<Type>`.
- Use **generic `Find` and `Traverse`** when working with **nested objects**.
- Full API documentation is available through GoDoc or your IDE.
- See [example](./example/) for practical usage.

## License

This package is licensed under [GNU-LGPL-3.0](./LICENSE).
