package transactions

import (
	"bankGolang/utils"
	"errors"
)

// #TODO remove the item list of brands, to be just a map

// Item struct for the item type
type Item struct {
	Name    string             `json:"name"`
	Brands  map[string]float64 `json:"brands"`
	Measure string             `json:"measure"`
}

// Message struct for the messages
type Message struct {
	Value string `json:"value"`
}

// #TODO RemoveBrand()

// AddBrand is a method that adds a brand to an item
func (item *Item) AddBrand(brands map[string]float64) {
	if item.Brands == nil {
		item.Brands = map[string]float64{}
	}
	for name, price := range brands {
		item.Brands[name] = float64(price)
	}
}

// #TODO Get the unit of an item from the web

// NewItem function creates a new item
func NewItem(name string, brands map[string]float64, unit string) (item *Item, err error) {
	if utils.CheckString(name) {
		item = &Item{}
		item.Name = name
		item.AddBrand(brands)
		if utils.CheckString(unit) {
			item.Measure = unit
			return
		}
		err = errors.New("invalid unit, empty string")
		return
	}
	err = errors.New("invalid name, empty string")
	return
}
