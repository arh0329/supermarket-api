package models

import "errors"

type allItems []Item

var errDuplicate = errors.New("product Code is already in use. Product Codes must be unique")

// Produce is the in memory database.
var Produce = allItems{
	{
		Name:        "Lettuce",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   3.46,
	},
	{
		Name:        "Peach",
		ProduceCode: "E5T6-9UI3-TH15-QR88",
		UnitPrice:   2.99,
	},
	{
		Name:        "Green Pepper",
		ProduceCode: "YRT6-72AS-K736-L4AR",
		UnitPrice:   0.79,
	},
	{
		Name:        "Gala Apple",
		ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
		UnitPrice:   3.59,
	},
}

// GetAllProduce retrieves all items from database. Returns an array of Items
func GetAllProduce() []Item {
	return Produce
}

// GetOneItem retrieves one item from the database. Accepts
// the product code of the item to be retrieved. Returns item if found.
// Returns empty item if not found
func GetOneItem(pc string) Item {
	empty := Item{}
	for _, item := range Produce {
		if item.ProduceCode == pc {
			return item
		}
	}
	return empty
}

// AddProduce adds item to database. Accepts item to be added.
// Returns error if item product code is already present in database
func AddProduce(item Item) error {
	if !isDuplicate(item) {
		Produce = append(Produce, item)
	} else {
		return errDuplicate
	}
	return nil
}

// DeleteProduce deletes an item from the database. Accepts the product code
// of the item to be deleted. Returns deleted item. Returns empty item if not found
func DeleteProduce(pc string) Item {
	empty := Item{}
	for i, item := range Produce {
		if item.ProduceCode == pc {
			Produce = append(Produce[:i], Produce[i+1:]...)
			return item
		}
	}
	return empty
}

// Equal checks if two arrays of Items are equal
func Equal(a, b []Item) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// isDuplicate checks if item is a already in the database
func isDuplicate(newItem Item) bool {

	for _, item := range Produce {
		if item.ProduceCode == newItem.ProduceCode {
			return true
		}
	}
	return false
}
