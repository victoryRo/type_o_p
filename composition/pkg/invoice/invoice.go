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
	items   invoiceitem.Items
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) *Invoice {
	return &Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
