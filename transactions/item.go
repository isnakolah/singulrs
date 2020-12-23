package transactions

import (
	"errors"
	"strings"
)

// #TODO remove the item list of brands, to be just a map

// Item struct for the item type
type Item struct {
	Name    string
	Brands  []map[string]float64
	Measure string
}

// Message struct for the messages
type Message struct {
	Value string
}

// #TODO RemoveBrand()

// AddBrand is a method that adds a brand to an item
func (item *Item) AddBrand(brands map[string]float64) {
	for name, price := range brands {
		item.Brands = append(item.Brands, map[string]float64{name: price})
	}
}

// #TODO Get the unit of a good from the web

// NewItem function creates a new item
func NewItem(name string, brands map[string]float64, unit string) (item *Item, err error) {
	if strings.TrimSpace(name) != "" {
		item = &Item{}
		item.Name = name
		item.AddBrand(brands)
		if strings.TrimSpace(unit) != "" {
			item.Measure = unit
			return
		}
		err = errors.New("invalid unit, empty string")
		return
	}
	err = errors.New("invalid name, empty string")
	return
}
