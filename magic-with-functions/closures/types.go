package closures

// IsBiggerThan returns a closure function that is parametrized with the
// first parameter
func IsBiggerThan(first int) func(second int) bool {
	isBigger := func(second int) bool {
		return second > first
	}

	return isBigger
}

type Customer struct {
	Name string
}

type CustomerID struct {
	Value string
}

type FileBasedCustomerRepository struct{}

func NewFileBasedCustomerRepository() *FileBasedCustomerRepository {
	return &FileBasedCustomerRepository{}
}

func (r *FileBasedCustomerRepository) Add(customer Customer) error {
	return nil
}

func (r *FileBasedCustomerRepository) Delete(id CustomerID) error {
	return nil
}

type CustomerRepository interface {
	Add(customer Customer) error
	Delete(id CustomerID) error
}

// CustomerService has repo as a protperty
// that will be "enclosed" in the receiver methods
type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

// Register is a receiver method that will be used as a closure
func (s *CustomerService) Register(customerName string) error {
	customer := Customer{Name: customerName}

	return s.repo.Add(customer)
}

// Remove is a receiver method that will be used as a closure
func (s *CustomerService) Remove(id CustomerID) error {
	return s.repo.Delete(id)
}

// ForRegisteringCustomers is a function type
type ForRegisteringCustomers func(customerName string) error

// ForRemovingCustomers is a function type
type ForRemovingCustomers func(id CustomerID) error

// CustomerRegistrationHTTPHandler receives a closure as dependency
// which fullfulls the function type ForRegisteringCustomers
type CustomerRegistrationHTTPHandler struct {
	registerCustomer ForRegisteringCustomers
}

func NewCustomerRegistrationHTTPHandler(
	registerCustomer ForRegisteringCustomers,
) *CustomerRegistrationHTTPHandler {

	return &CustomerRegistrationHTTPHandler{
		registerCustomer: registerCustomer,
	}
}

func (h *CustomerRegistrationHTTPHandler) Handle(customerName string) error {
	return h.registerCustomer(customerName)
}

// CustomerRemovalHTTPHandler receives a closure as dependency
// which fullfulls the function type ForRegisteringCustomers
type CustomerRemovalHTTPHandler struct {
	removeCustomer ForRemovingCustomers
}

func NewCustomerRemovalHTTPHandler(
	removeCustomer ForRemovingCustomers,
) *CustomerRemovalHTTPHandler {

	return &CustomerRemovalHTTPHandler{
		removeCustomer: removeCustomer,
	}
}

func (h *CustomerRemovalHTTPHandler) Remove(id CustomerID) error {
	return h.removeCustomer(id)
}
