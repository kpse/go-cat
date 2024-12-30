package functor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe(t *testing.T) {
	t.Run("Just operations", func(t *testing.T) {
		m := Just(5)
		assert.True(t, m.IsJust())
		assert.False(t, m.IsNothing())
		assert.Equal(t, 5, m.Get())
		assert.Equal(t, 5, m.GetOrElse(10))
	})

	t.Run("Nothing operations", func(t *testing.T) {
		m := Nothing[int]()
		assert.False(t, m.IsJust())
		assert.True(t, m.IsNothing())
		assert.Panics(t, func() { m.Get() })
		assert.Equal(t, 10, m.GetOrElse(10))
	})

	t.Run("Map", func(t *testing.T) {
		double := func(x int) int { return x * 2 }

		just := Map(Just(5), double)
		assert.Equal(t, 10, just.Get())

		nothing := Map(Nothing[int](), double)
		assert.True(t, nothing.IsNothing())
	})

	t.Run("FlatMap", func(t *testing.T) {
		safeDivide := func(x int) Maybe[int] {
			if x == 0 {
				return Nothing[int]()
			}
			return Just(100 / x)
		}

		just := FlatMap(Just(5), safeDivide)
		assert.Equal(t, 20, just.Get())

		nothing := FlatMap(Just(0), safeDivide)
		assert.True(t, nothing.IsNothing())
	})

	t.Run("Filter", func(t *testing.T) {
		isEven := func(x int) bool { return x%2 == 0 }

		even := Filter(Just(4), isEven)
		assert.True(t, even.IsJust())

		odd := Filter(Just(5), isEven)
		assert.True(t, odd.IsNothing())
	})
}
