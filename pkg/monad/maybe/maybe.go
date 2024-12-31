package maybe

// Maybe represents an optional value
type Maybe[T any] struct {
	value *T
}

func Just[T any](x T) Maybe[T] {
	return Maybe[T]{value: &x}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{value: nil}
}

func (m Maybe[T]) Get() T {
	if m.value == nil {
		panic("Called Get on Nothing")
	}
	return *m.value
}

func (m Maybe[T]) GetOrElse(defaultVal T) T {
	if m.value == nil {
		return defaultVal
	}
	return *m.value
}

func (m Maybe[T]) IsJust() bool {
	return m.value != nil
}

func (m Maybe[T]) IsNothing() bool {
	return m.value == nil
}

func Map[T, U any](m Maybe[T], f func(T) U) Maybe[U] {
	if m.value == nil {
		return Nothing[U]()
	}
	result := f(*m.value)
	return Just(result)
}

func FlatMap[T, U any](m Maybe[T], f func(T) Maybe[U]) Maybe[U] {
	if m.value == nil {
		return Nothing[U]()
	}
	return f(*m.value)
}

func Filter[T any](m Maybe[T], pred func(T) bool) Maybe[T] {
	if m.value == nil || !pred(*m.value) {
		return Nothing[T]()
	}
	return m
}
