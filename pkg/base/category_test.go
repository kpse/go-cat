package base_test

import (
	"testing"

	"github.com/kpse/go-cat/pkg/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCategory_Creation tests the basic creation and initialization of a category
func TestCategory_Creation(t *testing.T) {
	t.Run("new category should be empty", func(t *testing.T) {
		cat := base.NewCategory[int]()
		assert.Empty(t, cat.Objects)
		assert.Empty(t, cat.Morphisms)
	})
}

// TestCategory_AddObject tests adding objects to a category
func TestCategory_AddObject(t *testing.T) {
	t.Run("should add object and initialize morphism map", func(t *testing.T) {
		cat := base.NewCategory[int]()
		cat.AddObject(1)

		assert.Contains(t, cat.Objects, 1)
		assert.NotNil(t, cat.Morphisms[1])
	})
}

// TestCategory_AddMorphism tests adding and composing morphisms
func TestCategory_AddMorphism(t *testing.T) {
	t.Run("should add morphism between objects", func(t *testing.T) {
		cat := base.NewCategory[int]()
		cat.AddObject(1)
		cat.AddObject(2)

		double := func(x int) int { return x * 2 }
		cat.AddMorphism(1, 2, double)

		morphisms := cat.Morphisms[1][2]
		require.Len(t, morphisms, 1)
		assert.Equal(t, 2, morphisms[0].Transform(1))
	})
}

// TestMorphism_Composition tests morphism composition
func TestMorphism_Composition(t *testing.T) {
	t.Run("should correctly compose two morphisms", func(t *testing.T) {
		// Create morphisms f: A → B and g: B → C
		f := base.Morphism[int, int]{
			Source:    1,
			Target:    2,
			Transform: func(x int) int { return x * 2 },
		}

		g := base.Morphism[int, int]{
			Source:    2,
			Target:    3,
			Transform: func(x int) int { return x + 1 },
		}

		// Compose them to get h = g ∘ f
		h := base.Compose(f, g)

		// Test composition: h(1) should be (1 * 2) + 1 = 3
		assert.Equal(t, 3, h.Transform(1))
		assert.Equal(t, 1, h.Source)
		assert.Equal(t, 3, h.Target)
	})
}

// TestIdentity_Morphism tests identity morphism properties
func TestIdentity_Morphism(t *testing.T) {
	t.Run("identity morphism should preserve value", func(t *testing.T) {
		id := base.Identity(42)
		assert.Equal(t, 42, id.Transform(42))
	})

	t.Run("identity should satisfy left identity law", func(t *testing.T) {
		f := base.Morphism[int, int]{
			Source:    1,
			Target:    2,
			Transform: func(x int) int { return x * 2 },
		}

		idA := base.Identity(1)
		composed := base.Compose(idA, f)

		assert.Equal(t, f.Transform(5), composed.Transform(5))
	})

	t.Run("identity should satisfy right identity law", func(t *testing.T) {
		f := base.Morphism[int, int]{
			Source:    1,
			Target:    2,
			Transform: func(x int) int { return x * 2 },
		}

		idB := base.Identity(2)
		composed := base.Compose(f, idB)

		assert.Equal(t, f.Transform(5), composed.Transform(5))
	})
}

// TestCategory_Laws tests the category laws
func TestCategory_Laws(t *testing.T) {
	t.Run("composition should be associative", func(t *testing.T) {
		// Create three morphisms f: A → B, g: B → C, h: C → D
		f := base.Morphism[int, int]{
			Transform: func(x int) int { return x * 2 },
		}
		g := base.Morphism[int, int]{
			Transform: func(x int) int { return x + 1 },
		}
		h := base.Morphism[int, int]{
			Transform: func(x int) int { return x * 3 },
		}

		// Test (h ∘ g) ∘ f = h ∘ (g ∘ f)
		comp1 := base.Compose(base.Compose(f, g), h)
		comp2 := base.Compose(f, base.Compose(g, h))

		testValue := 5
		assert.Equal(t, comp1.Transform(testValue), comp2.Transform(testValue))
	})
}

// Benchmark_Morphism_Composition benchmarks morphism composition
func Benchmark_Morphism_Composition(b *testing.B) {
	f := base.Morphism[int, int]{
		Transform: func(x int) int { return x * 2 },
	}
	g := base.Morphism[int, int]{
		Transform: func(x int) int { return x + 1 },
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := base.Compose(f, g)
		_ = h.Transform(i)
	}
}
