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

type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Register(customerName string) error {
	customer := Customer{Name: customerName}

	return s.repo.Add(customer)
}

func (s *CustomerService) Remove(id CustomerID) error {
	return s.repo.Delete(id)
}

type ForRegisteringCustomers func(customerName string) error

type ForRemovingCustomers func(id CustomerID) error

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
