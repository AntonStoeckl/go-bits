package basics_test

import (
	"testing"

	"github.com/matryer/is"

	. "go-bits/magic-with-functions/basics"
)

// Using a regular function
func Test_IsBigger_1(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)

	// Act
	result := IsBigger(2, 1)

	// Assert
	shouldBe.True(result)
}

// Using a function as a value
func Test_isBigger_2(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	isBigger := IsBigger

	// Act
	result := isBigger(2, 1)

	// Assert
	shouldBe.True(result)
}

// Using a lambda function
func Test_isBigger_3(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	isBigger := func(first, second int) bool {
		return first > second
	}

	// Act
	result := isBigger(2, 1)

	// Assert
	shouldBe.True(result)
}

// Passing a regular function as a value to a higher order function
// with explicit annotation
func Test_Compare_1(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)

	// Act
	result := Compare(2, 1, IsBigger)

	// Assert
	shouldBe.True(result)
}

// Passing a lambda function as a value to a higher order function
// with explicit annotation
func Test_Compare_2(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	isBigger := func(first, second int) bool {
		return first > second
	}

	// Act
	result := Compare(2, 1, isBigger)

	// Assert
	shouldBe.True(result)
}

// Passing a regular function as a value to a higher order function
// with a function type annotation
func Test_OtherCompare1(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)

	// Act
	result := OtherCompare(2, 1, IsBigger)

	// Assert
	shouldBe.True(result)
}

// Passing a lambda function as a value to a higher order function
// with a function type annotation
func Test_OtherCompare_2(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	isBigger := func(first, second int) bool {
		return first > second
	}

	// Act
	result := OtherCompare(2, 1, isBigger)

	// Assert
	shouldBe.True(result)
}
