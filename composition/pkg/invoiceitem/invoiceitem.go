package invoiceitem

// Item contains information of an invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{
		id,
		product,
		value,
	}
}

// Items is a slice of Item
type Items []Item

func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}

	return is
}

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}

	return total
}
