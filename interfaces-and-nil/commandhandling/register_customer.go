package commandhandling

// RegisterCustomer implements the Command interface
type RegisterCustomer struct {
	commandType string
	// some useful properties
}

func NewRegisterCustomer() *RegisterCustomer {
	return &RegisterCustomer{
		commandType: "RegisterCustomer",
	}
}

func (c *RegisterCustomer) CommandType() string {
	return c.commandType
}
