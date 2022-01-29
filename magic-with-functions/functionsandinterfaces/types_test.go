package functionsandinterfaces_test

import (
	"testing"

	"github.com/matryer/is"

	. "go-bits/magic-with-functions/functionsandinterfaces"
)

func Test_isBigger_with_a_closure(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)

	realFunction := RegisterCustomer
	castedFunction := ForRegisteringCustomersFunc(realFunction)
	handler := NewCustomerRegistrationHTTPHandler(castedFunction)

	// Act
	err := handler.Handle("John Doe", "john@doe.com")

	// Assert
	shouldBe.NoErr(err)
}
