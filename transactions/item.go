package transactions

// Item struct for the item type
type Item struct {
	Name    string
	Brands  []map[string]float64
	Measure string
}

// Brand struct for the brand types
// type Brand struct {
// 	Name  string
// 	Price float64
// }

// AddBrand is a method that adds a brand to an item
func (item *Item) AddBrand(brands map[string]float64) {
	for name, price := range brands {
		item.Brands = append(item.Brands, map[string]float64{name: price})
	}
}

// NewItem function creates a new item
func NewItem(name string, brands map[string]float64, unit string) *Item {
	item := &Item{}
	item.Name = name
	item.Measure = unit
	item.AddBrand(brands)

	return item
}
