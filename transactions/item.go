package transactions

import (
	"fmt"
)

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

// NewItem method to add a new Item
func NewItem(name string, brands map[string]int) *Item {
	item := &Item{}
	item.Name = name

	for name, price := range brands {
		fmt.Println(name, price)
	}
	return item
}
