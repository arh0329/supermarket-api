package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_ValidItem(t *testing.T) {
	item := Item{
		Name:        "Banana",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   1.00,
	}
	err := item.Validate()
	assert.Nil(t, err)
}

func Test_Validate_BadName(t *testing.T) {
	item := Item{
		Name:        "Ban#na",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   1.00,
	}
	err := item.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "name is invalid. Must contain only alphanumeric characters and be between 2-50 characters in length", err.Error())
}

func Test_Validate_BadProductCode(t *testing.T) {
	item := Item{
		Name:        "Banana",
		ProduceCode: "A12T-4GH7-QPL9-3N4%M",
		UnitPrice:   1.00,
	}
	err := item.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, errInvalidProductCode, err)
}

func Test_Validate_NoUnitPrice(t *testing.T) {
	item := Item{
		Name:        "Banana",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
	}
	err := item.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Key: 'Item.UnitPrice' Error:Field validation for 'UnitPrice' failed on the 'required' tag", err.Error())
}

func Test_Validate_WillRoundUnitPriceToTwoDecimals(t *testing.T) {
	item := Item{
		Name:        "Banana",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   1.95678,
	}
	roundedItem := Item{
		Name:        "Banana",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   1.96,
	}
	err := item.Validate()
	assert.Nil(t, err)
	assert.Equal(t, roundedItem, item)
}
