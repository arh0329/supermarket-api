package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var produceStart = allItems{
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

var produceAlt = allItems{
	{
		Name:        "Mango",
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

var produceWithPumpkin = allItems{
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

var produceWithoutLettuce = allItems{
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

func resetDatabase() {
	Produce = produceStart
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
	resetDatabase()
	AddProduce(pumpkin)
	assert.Equal(t, produceWithPumpkin, Produce)
}

func Test_AddProduce_ReturnsDuplicateErr(t *testing.T) {
	resetDatabase()
	_ = AddProduce(pumpkin)
	err := AddProduce(pumpkin)
	assert.Equal(t, errDuplicate, err)
	assert.Equal(t, produceWithPumpkin, Produce)
}

func Test_DeleteProduce_DeletesItem(t *testing.T) {
	resetDatabase()
	_ = DeleteProduce("A12T-4GH7-QPL9-3N4M")
	assert.Equal(t, produceWithoutLettuce, Produce)
}

func Test_DeleteProduce_ReturnsEmptyItem(t *testing.T) {
	result := DeleteProduce("A12T-4GH7-QPL9-3N40")
	empty := Item{}
	assert.Equal(t, empty, result)
}

func Test_Equal_ReturnsFalse_IfArraysNotEqualLength(t *testing.T) {
	array1 := produceStart
	array2 := produceWithPumpkin
	res := Equal(array1, array2)
	assert.False(t, res)
}

func Test_Equal_ReturnsFalse_IfArraysNotEqual(t *testing.T) {
	array1 := produceStart
	array2 := produceAlt
	res := Equal(array1, array2)
	assert.False(t, res)
}

func Test_Equal_ReturnsTrue_IfArraysEqual(t *testing.T) {
	array1 := produceStart
	array2 := produceStart
	res := Equal(array1, array2)
	assert.True(t, res)
}

func Test_isDuplicate_ReturnsTrue_IfDuplicate(t *testing.T) {
	res := isDuplicate(Produce[0])
	assert.True(t, res)
}

func Test_isDuplicate_ReturnsFalse_IfNotDuplicate(t *testing.T) {
	res := isDuplicate(pumpkin)
	assert.False(t, res)
}
