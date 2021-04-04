package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

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

var ProduceWithPumpkin = allItems{
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
	{
		Name:        "Pumpkin",
		ProduceCode: "TQ4C-VV6T-75ZX-1RM4",
		UnitPrice:   5.67,
	},
}

var ProduceWithoutLettuce = allItems{
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

var pumpkin = Item{
	Name:        "Pumpkin",
	ProduceCode: "TQ4C-VV6T-75ZX-1RM4",
	UnitPrice:   5.67,
}

func Test_GetAllProduce(t *testing.T) {
	result := GetAllProduce()
	if !Equal(Produce, result) {
		t.Errorf("Actual does not equal expected.\nActual:%v\nExpected:%v", Produce, result)
	}
}

func Test_GetOneItem_ReturnsItem(t *testing.T) {
	result := GetOneItem("E5T6-9UI3-TH15-QR88")
	assert.Equal(t, result, Produce[1])
}

func Test_GetOneItem_ReturnsEmptyItem(t *testing.T) {
	result := GetOneItem("E5T6-9UI3-TH15-QR89")
	empty := Item{}
	assert.Equal(t, empty, result)
}

func Test_AddProduce_AddsItemToDatabase(t *testing.T) {
	AddProduce(pumpkin)
	assert.Equal(t, ProduceWithPumpkin, Produce)
}

func Test_DeleteProduce_DeletesItem(t *testing.T) {
	_ = DeleteProduce("A12T-4GH7-QPL9-3N4M")
	assert.Equal(t, ProduceWithoutLettuce, Produce)
}

func Test_DeleteProduce_ReturnsEmptyItem(t *testing.T) {
	result := DeleteProduce("A12T-4GH7-QPL9-3N40")
	empty := Item{}
	assert.Equal(t, empty, result)
}
