package closures_test

import (
	"testing"

	"github.com/matryer/is"

	. "go-bits/magic-with-functions/closures"
)

// Build a parametrized function with a closure
func Test_isBigger_with_a_closure(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	isBiggerThan1 := IsBiggerThan(1)

	// Act
	result := isBiggerThan1(2)

	// Assert
	shouldBe.True(result)
}

func Test_object_with_receiver_method_as_closure(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	repo := NewFileBasedCustomerRepository()
	service := NewCustomerService(repo)
	registrationHandler := NewCustomerRegistrationHTTPHandler(service.Register)
	removalHandler := NewCustomerRemovalHTTPHandler(service.Remove)

	// Act
	registerErr := registrationHandler.Handle("John Doe")

	// Assert
	shouldBe.NoErr(registerErr)

	// Clean up
	removalErr := removalHandler.Remove(CustomerID{Value: "123435"})
	shouldBe.NoErr(removalErr)
}
