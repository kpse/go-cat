package base

import (
	"fmt"
	"testing"
)

func TestEitherConstructors(t *testing.T) {
	// Test Left constructor
	left := Left[string, int]("error")
	if !left.IsLeft() {
		t.Error("Left constructor created Right value")
	}
	if left.GetLeft() != "error" {
		t.Error("Left value not stored correctly")
	}

	// Test Right constructor
	right := Right[string, int](42)
	if !right.IsRight() {
		t.Error("Right constructor created Left value")
	}
	if right.GetRight() != 42 {
		t.Error("Right value not stored correctly")
	}
}

func TestEitherPanics(t *testing.T) {
	// Test GetLeft on Right value
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetLeft on Right value did not panic")
		}
	}()
	right := Right[string, int](42)
	_ = right.GetLeft()
}

func TestEitherGetRightPanic(t *testing.T) {
	// Test GetRight on Left value
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetRight on Left value did not panic")
		}
	}()
	left := Left[string, int]("error")
	_ = left.GetRight()
}

func TestMap(t *testing.T) {
	// Test Map on Right value
	right := Right[string, int](42)
	doubled := Map(right, func(x int) int { return x * 2 })
	if doubled.IsLeft() {
		t.Error("Map on Right value returned Left")
	}
	if doubled.GetRight() != 84 {
		t.Error("Map function not applied correctly")
	}

	// Test Map on Left value
	left := Left[string, int]("error")
	mappedLeft := Map(left, func(x int) int { return x * 2 })
	if !mappedLeft.IsLeft() {
		t.Error("Map on Left value returned Right")
	}
	if mappedLeft.GetLeft() != "error" {
		t.Error("Left value modified by Map")
	}
}

func TestBind(t *testing.T) {
	// Helper function that returns Either
	divide := func(x int) Either[string, float64] {
		if x == 0 {
			return Left[string, float64]("division by zero")
		}
		return Right[string, float64](100.0 / float64(x))
	}

	// Test Bind on Right value
	right := Right[string, int](50)
	result := Bind(right, divide)
	if result.IsLeft() {
		t.Error("Bind on valid division returned Left")
	}
	if result.GetRight() != 2.0 {
		t.Error("Bind computation incorrect")
	}

	// Test Bind on Right value that produces Left
	rightZero := Right[string, int](0)
	resultZero := Bind(rightZero, divide)
	if !resultZero.IsLeft() {
		t.Error("Bind on division by zero returned Right")
	}
	if resultZero.GetLeft() != "division by zero" {
		t.Error("Bind error message incorrect")
	}

	// Test Bind on Left value
	left := Left[string, int]("original error")
	resultLeft := Bind(left, divide)
	if !resultLeft.IsLeft() {
		t.Error("Bind on Left value returned Right")
	}
	if resultLeft.GetLeft() != "original error" {
		t.Error("Left value modified by Bind")
	}
}

func TestMatch(t *testing.T) {
	// Test Match on Right value
	right := Right[string, int](42)
	rightResult := Match(right,
		func(e string) string { return fmt.Sprintf("Error: %s", e) },
		func(x int) string { return fmt.Sprintf("Success: %d", x) },
	)
	if rightResult != "Success: 42" {
		t.Error("Match on Right value returned incorrect result")
	}

	// Test Match on Left value
	left := Left[string, int]("error")
	leftResult := Match(left,
		func(e string) string { return fmt.Sprintf("Error: %s", e) },
		func(x int) string { return fmt.Sprintf("Success: %d", x) },
	)
	if leftResult != "Error: error" {
		t.Error("Match on Left value returned incorrect result")
	}
}

func TestFromNillable(t *testing.T) {
	// Test with non-nil value
	nonNil := FromNillable(42, "error")
	if !nonNil.IsRight() {
		t.Error("FromNillable with non-nil value returned Left")
	}
	if nonNil.GetRight() != 42 {
		t.Error("FromNillable stored incorrect value")
	}

	// Test with zero/nil value
	var zero int
	nil := FromNillable(zero, "error")
	if !nil.IsLeft() {
		t.Error("FromNillable with nil value returned Right")
	}
	if nil.GetLeft() != "error" {
		t.Error("FromNillable stored incorrect error")
	}
}

func TestBiMap(t *testing.T) {
	// Test BiMap on Right value
	right := Right[string, int](42)
	mappedRight := BiMap(right,
		func(e string) int { return len(e) },
		func(x int) string { return fmt.Sprintf("%d", x*2) },
	)
	if !mappedRight.IsRight() {
		t.Error("BiMap on Right value returned Left")
	}
	if mappedRight.GetRight() != "84" {
		t.Error("BiMap right function not applied correctly")
	}

	// Test BiMap on Left value
	left := Left[string, int]("error")
	mappedLeft := BiMap(left,
		func(e string) int { return len(e) },
		func(x int) string { return fmt.Sprintf("%d", x*2) },
	)
	if !mappedLeft.IsLeft() {
		t.Error("BiMap on Left value returned Right")
	}
	if mappedLeft.GetLeft() != 5 {
		t.Error("BiMap left function not applied correctly")
	}
}

func TestFold(t *testing.T) {
	// Test Fold on Right value
	right := Right[string, int](42)
	rightResult := Fold(right, 10, func(acc int, x int) int { return acc + x })
	if rightResult != 52 {
		t.Error("Fold on Right value returned incorrect result")
	}

	// Test Fold on Left value
	left := Left[string, int]("error")
	leftResult := Fold(left, 10, func(acc int, x int) int { return acc + x })
	if leftResult != 10 {
		t.Error("Fold on Left value didn't return initial value")
	}
}
