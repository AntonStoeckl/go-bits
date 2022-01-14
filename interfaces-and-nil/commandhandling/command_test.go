package commandhandling_test

import (
	"testing"

	"go-bits/interfaces-and-nil/commandhandling"
)

func TestNilInterface(_ *testing.T) {
	var nilInterface commandhandling.Command

	if nilInterface != nil {
		_ = nilInterface.CommandType()
	}
}

func TestNilPointerCommand(_ *testing.T) {
	var nilCommand *commandhandling.RegisterCustomer

	if nilCommand != nil {
		_ = nilCommand.CommandType()
	}
}

func TestInterfaceWithUnderlyingNilPointer(_ *testing.T) {
	var nilCommand *commandhandling.RegisterCustomer
	var interfaceWithUnderlyingNilPointer commandhandling.Command
	interfaceWithUnderlyingNilPointer = nilCommand

	if interfaceWithUnderlyingNilPointer != nil {
		_ = interfaceWithUnderlyingNilPointer.CommandType()
	}
}
