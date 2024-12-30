package testutil

import (
	"testing"

	"github.com/kpse/go-cat/pkg/base"
)

// MorphismMatcher is a test helper to verify morphism properties
type MorphismMatcher struct {
	t *testing.T
}

// NewMorphismMatcher creates a new morphism matcher
func NewMorphismMatcher(t *testing.T) *MorphismMatcher {
	return &MorphismMatcher{t: t}
}

// Instead of methods, we'll use package-level functions
// AssertComposition verifies that morphism composition is correct
func AssertComposition[A, B, C comparable](
	t *testing.T,
	f base.Morphism[A, B],
	g base.Morphism[B, C],
	input A,
	expected C,
) {
	composed := base.Compose(f, g)
	result := composed.Transform(input)
	if result != expected {
		t.Errorf("Composition failed. Got %v, expected %v", result, expected)
	}
}

// AssertIdentityLaws verifies that a morphism satisfies identity laws
func AssertIdentityLaws[A, B comparable](
	t *testing.T,
	f base.Morphism[A, B],
	input A,
) {
	idA := base.Identity(f.Source)
	idB := base.Identity(f.Target)

	// Left identity: id ∘ f = f
	leftComp := base.Compose(idA, f)

	// Right identity: f ∘ id = f
	rightComp := base.Compose(f, idB)

	expected := f.Transform(input)
	if leftComp.Transform(input) != expected {
		t.Error("Left identity law failed")
	}
	if rightComp.Transform(input) != expected {
		t.Error("Right identity law failed")
	}
}

// TestCategory provides a simple category for testing
type TestCategory struct {
	base.Category[int]
}

// NewTestCategory creates a new test category with some common morphisms
func NewTestCategory() *TestCategory {
	cat := &TestCategory{
		Category: *base.NewCategory[int](),
	}
	cat.AddObject(1)
	cat.AddObject(2)
	cat.AddObject(3)

	// Add some basic morphisms
	cat.AddMorphism(1, 2, func(x int) int { return x * 2 })
	cat.AddMorphism(2, 3, func(x int) int { return x + 1 })

	return cat
}
