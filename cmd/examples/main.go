package main

import (
	"fmt"

	"github.com/kpse/go-cat/pkg/base"
)

func main() {
	// Create a category of integers
	cat := base.NewCategory[int]()

	// Add objects
	cat.AddObject(1)
	cat.AddObject(2)
	cat.AddObject(4)

	// Define morphisms
	double := func(x int) int { return x * 2 }
	cat.AddMorphism(1, 2, double)
	cat.AddMorphism(2, 4, double)

	// Create and compose morphisms
	f := base.Morphism[int, int]{
		Source:    1,
		Target:    2,
		Transform: double,
	}

	g := base.Morphism[int, int]{
		Source:    2,
		Target:    4,
		Transform: double,
	}

	// Compose f and g
	h := base.Compose(f, g)

	// Test the composition
	input := 1
	fmt.Printf("Input: %d\n", input)
	fmt.Printf("f(input): %d\n", f.Transform(input))
	fmt.Printf("g(f(input)): %d\n", h.Transform(input))

	// Test identity
	id := base.Identity(input)
	fmt.Printf("identity(input): %d\n", id.Transform(input))
}
