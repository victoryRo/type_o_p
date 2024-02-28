package customer

// Customer is the struct of a client
type Customer struct {
	name    string
	address string
	phone   string
}

// New creates a new customer with the given name, address, and phone number.
func New(name, address, phone string) Customer {
	return Customer{
		name,
		address,
		phone,
	}
}
