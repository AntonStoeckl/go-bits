package commandhandling

import (
	"errors"
)

var ErrCommandIsNilInterface = errors.New("command is nil interface")
var ErrCommandIsNilPointer = errors.New("command is nil pointer")
var ErrUnknownCommand = errors.New("unknown command received")

type CommandHandler struct{}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

func (h *CommandHandler) Handle(command Command) error {
	var err error

	switch actual := command.(type) {
	case *RegisterCustomer:
		err = h.handleRegisterCustomer(actual)
	case *ConfirmEmailAddress:
		err = h.handleConfirmEmailAddress(actual)
	case nil:
		return ErrCommandIsNilInterface
	default:
		return ErrUnknownCommand
	}

	return err
}

func (h *CommandHandler) handleRegisterCustomer(command *RegisterCustomer) error {
	if command == nil {
		return ErrCommandIsNilPointer
	}

	// do something useful
	_ = command.CommandType()

	return nil
}

func (h *CommandHandler) handleConfirmEmailAddress(command *ConfirmEmailAddress) error {
	if command == nil {
		return ErrCommandIsNilPointer
	}

	// do something useful
	_ = command.CommandType()

	return nil
}
