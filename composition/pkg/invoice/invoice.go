package invoice

import (
	"github.com/victoryRo/type_o_p/composition/pkg/customer"
	"github.com/victoryRo/type_o_p/composition/pkg/invoiceitem"
)

// Invoice is the struct of an invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) *Invoice {
	return &Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal sets the Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
