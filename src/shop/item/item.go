package item

import (
	"errors"
	"fmt"
	"singulr/src/utils/checkstring"
)

// #TODO remove the item list of brands, to be just a map

// Item struct for the item type
type Item struct {
	Name   string             `json:"name"`
	Brands map[string]float64 `json:"brands"`
	Unit   string             `json:"unit"`
}

// Message struct for the messages
type Message struct {
	Value string `json:"value"`
}

// AddBrand is a method that adds a brand to an item
func (item *Item) AddBrand(brand map[string]float64) (message string) {
	if item.Brands == nil {
		item.Brands = map[string]float64{}
	}
	for name, price := range brand {
		findBrand := func() (available bool) {
			if _, ok := item.Brands[name]; ok {
				available = true
			}
			return
		}

		if !findBrand() {
			item.Brands[name] = float64(price)
			message = fmt.Sprintf("%s brand added.", name)
		}
	}
	return
}

// #TODO Get the unit of an item from the web

// New function creates a new item
func New(name string, brands map[string]float64, unit string) (item *Item, err error) {
	if checkstring.CheckString(name) {
		item = &Item{}
		item.Name = name
		item.AddBrand(brands)
		if checkstring.CheckString(unit) {
			item.Unit = unit
			return
		}
		err = errors.New("invalid unit, empty string")
		return
	}
	err = errors.New("invalid name, empty string")
	return
}
