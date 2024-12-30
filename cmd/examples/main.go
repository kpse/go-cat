package main

import (
	"fmt"

	"github.com/kpse/go-cat/pkg/base"
	"github.com/kpse/go-cat/pkg/functor"
)

func safeDivide(x int) functor.Maybe[int] {
	if x == 0 {
		return functor.Nothing[int]()
	}
	return functor.Just(100 / x)
}

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

	result := functor.Just(20)
	result = functor.Map(result, double)
	result = functor.FlatMap(result, safeDivide)
	result = functor.Filter(result, func(x int) bool { return x > 0 })

	fmt.Printf("Result: %v\n", result.Get())

	// Chain operations with Nothing
	noResult := functor.Just(0)
	noResult = functor.Map(noResult, double)
	noResult = functor.FlatMap(noResult, safeDivide)
	noResult = functor.Filter(noResult, func(x int) bool { return x > 0 })

	fmt.Printf("Has result: %v\n", noResult.IsJust())
}
