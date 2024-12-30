# go-cat

A Go library implementing category theory concepts and abstractions.

## Overview

go-cat provides practical implementations of category theory concepts in Go, including:

- Categories
- Functors
- Monads
- Natural Transformations

## Installation

```bash
go get github.com/kpse/go-cat
```

## Quick Start

Here's a simple example using the Maybe functor:

```go
package main

import (
    "fmt"
    "github.com/kpse/go-cat/pkg/functor"
)

func main() {
    // Create a Maybe functor
    maybeInt := functor.NewMaybe(5)

    // Map a function over the Maybe
    doubled := maybeInt.Map(func(x int) int {
        return x * 2
    })

    fmt.Printf("Result: %v\n", doubled.Get())
}
```

## Features

- Type-safe implementations using Go generics
- Practical abstractions for common programming patterns
- Extensive documentation and examples
- Full test coverage

## Documentation

For detailed documentation, please visit our [documentation](./docs/README.md).

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](./docs/CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
