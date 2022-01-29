package functionsandinterfaces

type ForRegisteringCustomers interface {
	Register(name, emailAddress string) error
}

type ForRegisteringCustomersFunc func(name, emailAddress string) error

// Register is a receiver method for the ForRegisteringCustomersFunc
// type that fullfills the ForRegisteringCustomers interface
func (f ForRegisteringCustomersFunc) Register(
	name string,
	emailAddress string,
) error {

	return f(name, emailAddress)
}

// RegisterCustomer is just a plain function that can be type-casted
// to ForRegisteringCustomersFunc
func RegisterCustomer(name, emailAddress string) error {
	// do something useful
	_, _ = name, emailAddress

	return nil
}

// CustomerRegistrationHTTPHandler has a dependency to something
// that fulfills the ForRegisteringCustomers interface
type CustomerRegistrationHTTPHandler struct {
	forRegisteringCustomers ForRegisteringCustomers
}

// NewCustomerRegistrationHTTPHandler is just a factory method
func NewCustomerRegistrationHTTPHandler(
	forRegisteringCustomers ForRegisteringCustomers,
) *CustomerRegistrationHTTPHandler {

	return &CustomerRegistrationHTTPHandler{
		forRegisteringCustomers: forRegisteringCustomers,
	}
}

// Handle uses the ForRegisteringCustomers dependency
func (h *CustomerRegistrationHTTPHandler) Handle(
	name string,
	emailAddress string,
) error {

	return h.forRegisteringCustomers.Register(name, emailAddress)
}
