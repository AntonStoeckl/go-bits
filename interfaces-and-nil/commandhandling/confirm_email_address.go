package commandhandling

// ConfirmEmailAddress implements the Command interface
type ConfirmEmailAddress struct {
	commandType string
	// some useful properties
}

func NewConfirmEmailAddress() *ConfirmEmailAddress {
	return &ConfirmEmailAddress{
		commandType: "ConfirmEmailAddress",
	}
}

func (c *ConfirmEmailAddress) CommandType() string {
	return c.commandType
}
