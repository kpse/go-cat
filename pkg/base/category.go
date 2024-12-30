package base

// Object represents any type that can be an object in a category
type Object interface{}

// Morphism represents a function between objects
type Morphism[A, B any] struct {
	Source    A
	Target    B
	Transform func(A) B
}

// Category represents a mathematical category
type Category[T comparable] struct {
	Objects   []T
	Morphisms map[T]map[T][]Morphism[T, T]
}

// NewCategory creates a new category
func NewCategory[T comparable]() *Category[T] {
	return &Category[T]{
		Objects:   make([]T, 0),
		Morphisms: make(map[T]map[T][]Morphism[T, T]),
	}
}

// AddObject adds an object to the category
func (c *Category[T]) AddObject(obj T) {
	c.Objects = append(c.Objects, obj)
	if _, exists := c.Morphisms[obj]; !exists {
		c.Morphisms[obj] = make(map[T][]Morphism[T, T])
	}
}

// AddMorphism adds a morphism between two objects
func (c *Category[T]) AddMorphism(source, target T, transform func(T) T) {
	morphism := Morphism[T, T]{
		Source:    source,
		Target:    target,
		Transform: transform,
	}

	if _, exists := c.Morphisms[source]; !exists {
		c.Morphisms[source] = make(map[T][]Morphism[T, T])
	}
	c.Morphisms[source][target] = append(c.Morphisms[source][target], morphism)
}

// Compose composes two compatible morphisms
func Compose[A, B, C any](f Morphism[A, B], g Morphism[B, C]) Morphism[A, C] {
	return Morphism[A, C]{
		Source: f.Source,
		Target: g.Target,
		Transform: func(a A) C {
			return g.Transform(f.Transform(a))
		},
	}
}

// Identity creates an identity morphism for an object
func Identity[T any](obj T) Morphism[T, T] {
	return Morphism[T, T]{
		Source:    obj,
		Target:    obj,
		Transform: func(t T) T { return t },
	}
}
