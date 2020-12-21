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

// New method to add a new Item
func (item Item) New(name string, brands map[string]int) Item {

}
