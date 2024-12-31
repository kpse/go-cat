package either

// Either represents a value of one of two possible types (a disjoint union)
type Either[E any, A any] struct {
	isLeft bool
	Left   E
	Right  A
}

// Left creates a new Either with a Left value
func Left[E any, A any](e E) Either[E, A] {
	return Either[E, A]{
		isLeft: true,
		Left:   e,
	}
}

// Right creates a new Either with a Right value
func Right[E any, A any](a A) Either[E, A] {
	return Either[E, A]{
		isLeft: false,
		Right:  a,
	}
}

// IsLeft returns true if the Either is a Left
func (e Either[E, A]) IsLeft() bool {
	return e.isLeft
}

// IsRight returns true if the Either is a Right
func (e Either[E, A]) IsRight() bool {
	return !e.isLeft
}

// GetLeft returns the Left value
// Panics if the Either is a Right
func (e Either[E, A]) GetLeft() E {
	if !e.isLeft {
		panic("Called GetLeft on Right value")
	}
	return e.Left
}

// GetRight returns the Right value
// Panics if the Either is a Left
func (e Either[E, A]) GetRight() A {
	if e.isLeft {
		panic("Called GetRight on Left value")
	}
	return e.Right
}

// Map applies a function to the Right value of an Either
func Map[E any, A any, B any](e Either[E, A], f func(A) B) Either[E, B] {
	if e.isLeft {
		return Left[E, B](e.Left)
	}
	return Right[E, B](f(e.Right))
}

// Bind chains computations with possible failures
func Bind[E any, A any, B any](e Either[E, A], f func(A) Either[E, B]) Either[E, B] {
	if e.isLeft {
		return Left[E, B](e.Left)
	}
	return f(e.Right)
}

// Match pattern matches on Either, applying the appropriate function
func Match[E any, A any, B any](e Either[E, A], leftFn func(E) B, rightFn func(A) B) B {
	if e.isLeft {
		return leftFn(e.Left)
	}
	return rightFn(e.Right)
}

// FromNillable creates an Either from a potentially nil value
func FromNillable[E any, A comparable](value A, err E) Either[E, A] {
	var zero A
	if value == zero {
		return Left[E, A](err)
	}
	return Right[E, A](value)
}

// BiMap applies functions to both sides of the Either
func BiMap[E1 any, E2 any, A1 any, A2 any](e Either[E1, A1], leftFn func(E1) E2, rightFn func(A1) A2) Either[E2, A2] {
	if e.isLeft {
		return Left[E2, A2](leftFn(e.Left))
	}
	return Right[E2, A2](rightFn(e.Right))
}

// Fold reduces the Either to a single value
func Fold[E any, A any, B any](e Either[E, A], initial B, f func(B, A) B) B {
	if e.isLeft {
		return initial
	}
	return f(initial, e.Right)
}
