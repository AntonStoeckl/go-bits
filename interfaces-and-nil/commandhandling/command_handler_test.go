package commandhandling_test

import (
	"testing"

	"github.com/matryer/is"

	"go-bits/interfaces-and-nil/commandhandling"
)

func TestHandlingTypedNilInterface(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	handler := commandhandling.NewCommandHandler()
	var nilInterface commandhandling.Command

	// Act
	err := handler.Handle(nilInterface)

	// Assert
	shouldBe.True(err != nil)
	shouldBe.True(err == commandhandling.ErrCommandIsNilInterface)
}

func TestHandlingUntypedNil(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	handler := commandhandling.NewCommandHandler()

	// Act
	err := handler.Handle(nil)

	// Assert
	shouldBe.True(err != nil)
	shouldBe.True(err == commandhandling.ErrCommandIsNilInterface)
}

func TestHandlingNilPointerCommand(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	handler := commandhandling.NewCommandHandler()
	var nilCommand *commandhandling.RegisterCustomer

	// Act
	err := handler.Handle(nilCommand)

	// Assert
	shouldBe.True(err != nil)
	shouldBe.True(err == commandhandling.ErrCommandIsNilPointer)
}

func TestHandlingInterfaceWithUnderlyingNilPointer(t *testing.T) {
	// Arrange
	shouldBe := is.New(t)
	handler := commandhandling.NewCommandHandler()
	var nilCommand *commandhandling.RegisterCustomer
	var interfaceWithUnderlyingNilPointer commandhandling.Command = nilCommand

	// Act
	err := handler.Handle(interfaceWithUnderlyingNilPointer)

	// Assert
	shouldBe.True(err != nil)
	shouldBe.True(err == commandhandling.ErrCommandIsNilPointer)
}
