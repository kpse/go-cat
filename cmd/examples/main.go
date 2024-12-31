package main

import (
	"fmt"
	"strconv"

	"github.com/kpse/go-cat/pkg/base"

	e "github.com/kpse/go-cat/pkg/monad/either"
	m "github.com/kpse/go-cat/pkg/monad/maybe"
)

func safeDivide(x int) m.Maybe[int] {
	if x == 0 {
		return m.Nothing[int]()
	}
	return m.Just(100 / x)
}

func divideBy(x, y int) e.Either[string, int] {
	if y == 0 {
		return e.Left[string, int]("division by zero")
	}
	return e.Right[string, int](x / y)
}

func parseInt(s string) e.Either[string, int] {
	if n, err := strconv.Atoi(s); err != nil {
		return e.Left[string, int]("parse error: " + err.Error())
	} else {
		return e.Right[string, int](n)
	}
}

func double(x int) int { return x * 2 }

func main() {
	basicCategoryCompose()

	maybeExamples()

	eitherExamples()
}

func eitherExamples() {
	fmt.Println("\n=== Basic Either Example ===")
	result2 := divideBy(10, 2)
	if result2.IsRight() {
		fmt.Printf("10/2 = %d\n", result2.GetRight())
	}

	errResult := divideBy(10, 0)
	if errResult.IsLeft() {
		fmt.Printf("Error: %s\n", errResult.GetLeft())
	}

	fmt.Println("\n=== Chaining Computations ===")

	chainResult := e.Map(parseInt("42"), func(x int) any {
		return divideBy(x, 2)
	})

	if chainResult.IsRight() {
		fmt.Printf("42/2 = %v\n", chainResult.GetRight().(e.Either[string, int]).GetRight())
	}

	badChain := e.Map(parseInt("not a number"), func(x int) any {
		return divideBy(x, 2)
	})

	if badChain.IsLeft() {
		fmt.Printf("Error: %s\n", badChain.GetLeft())
	}

	fmt.Println("\n=== FromNillable Example ===")
	var ptr *int
	nilResult := e.FromNillable(ptr, "nil pointer error")
	if nilResult.IsLeft() {
		fmt.Printf("Error: %s\n", nilResult.GetLeft())
	}

	value := 42
	ptrResult := e.FromNillable(&value, "nil pointer error")
	if ptrResult.IsRight() {
		fmt.Printf("Value: %d\n", ptrResult.GetRight())
	}

	fmt.Println("\n=== Map Transformation Example ===")
	numResult := e.Right[string, int](21)
	doubled := e.Map(numResult, func(x int) any {
		return x * 2
	})

	if doubled.IsRight() {
		fmt.Printf("21 * 2 = %v\n", doubled.GetRight())
	}

	fmt.Println("\n=== Complex Chain Example ===")
	complexChain := e.Map(parseInt("63"), func(x int) any {
		divided := divideBy(x, 3)
		if divided.IsLeft() {
			return divided
		}
		return e.Map(divided, func(y int) any {
			return y * 2
		})
	})

	if complexChain.IsRight() {
		result := complexChain.GetRight().(e.Either[string, any])
		if result.IsRight() {
			fmt.Printf("(63 / 3) * 2 = %v\n", result.GetRight())
		}
	}
}

func maybeExamples() {
	result := m.Just(20)
	result = m.Map(result, double)
	result = m.FlatMap(result, safeDivide)
	result = m.Filter(result, func(x int) bool { return x > 0 })

	fmt.Printf("Result: %v\n", result.Get())

	noResult := m.Just(0)
	noResult = m.Map(noResult, double)
	noResult = m.FlatMap(noResult, safeDivide)
	noResult = m.Filter(noResult, func(x int) bool { return x > 0 })

	fmt.Printf("Has result: %v\n", noResult.IsJust())
}

func basicCategoryCompose() {
	cat := base.NewCategory[int]()

	cat.AddObject(1)
	cat.AddObject(2)
	cat.AddObject(4)

	cat.AddMorphism(1, 2, double)
	cat.AddMorphism(2, 4, double)

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

	h := base.Compose(f, g)

	input := 1
	fmt.Printf("Input: %d\n", input)
	fmt.Printf("f(input): %d\n", f.Transform(input))
	fmt.Printf("g(f(input)): %d\n", h.Transform(input))

	id := base.Identity(input)
	fmt.Printf("identity(input): %d\n", id.Transform(input))
}
