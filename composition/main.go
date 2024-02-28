package main

import (
	"fmt"
	"github.com/victoryRo/type_o_p/composition/pkg/customer"
	"github.com/victoryRo/type_o_p/composition/pkg/invoice"
	"github.com/victoryRo/type_o_p/composition/pkg/invoiceitem"
)

func main() {
	client := customer.New("Victor", "Calle 19 la fortune", "664 52 15")
	items := []invoiceitem.Item{
		invoiceitem.New(1, "Laptops", 100.00),
		invoiceitem.New(2, "Computer", 40.00),
		invoiceitem.New(3, "Table", 80.10),
	}

	inv := invoice.New("Colombia", "Popayàn", client, items)
	inv.SetTotal()

	fmt.Printf("Invoice %+v\n", inv)
}
