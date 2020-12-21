package transactions

// Item struct for the item type
type Item struct {
	Name   string
	Brands []Brand
}

// Brand struct for the brand types
type Brand struct {
	Name  string
	Price float64
}

// AddBrand is a method that adds a brand to an item
func (item *Item) AddBrand(brands map[string]float64) {
	for name, price := range brands {
		item.Brands = append(item.Brands, Brand{name, price})
	}
}

// NewItem function creates a new item
func NewItem(name string, brands map[string]float64) *Item {
	item := &Item{}
	item.Name = name
	item.AddBrand(brands)

	return item
}
