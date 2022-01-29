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
	// ARRANGE
	shouldBe := is.New(t)

	repo := NewFileBasedCustomerRepository()

	// the repo from above is a property of the CustomerService struct
	// which is "enclosed" in the method closures below
	service := NewCustomerService(repo)

	// this is a receiver method used as a value
	register := service.Register

	// we are passing this method as a closure
	registrationHandler := NewCustomerRegistrationHTTPHandler(register)

	// this is a receiver method used as a value
	remove := service.Remove

	// we are passing this method as a closure
	removalHandler := NewCustomerRemovalHTTPHandler(remove)

	// ACT
	registerErr := registrationHandler.Handle("John Doe")

	// ASSERT
	shouldBe.NoErr(registerErr)

	// CLEAN UP
	removalErr := removalHandler.Remove(CustomerID{Value: "123435"})
	shouldBe.NoErr(removalErr)
}
