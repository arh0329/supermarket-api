package models

import (
	"errors"
	"math"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	errInvalidProductCode = errors.New("productCode is invalid must be only alphanumeric characters. The produce codes are sixteen characters long, with dashes separating each four character group")
	errInvalidName        = errors.New("name is invalid. Must contain only alphanumeric characters and be between 2-50 characters in length")
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Item struct {
	Name        string  `json:"name" validate:"required"`
	ProduceCode string  `json:"produceCode" validate:"required"`
	UnitPrice   float64 `json:"unitPrice" validate:"required,numeric"`
}

// Validate checks Item to ensure it follows conventions for name and product code
// It also rounds unitPrice to two decimals
func (i *Item) Validate() (err error) {
	err = validate.Struct(i)
	if err != nil {
		return err
	}

	if i.Name != "" {
		if valid, _ := regexp.Match(`^[a-zA-Z0-9 ]{2,50}$`, []byte(i.Name)); !valid {
			return errInvalidName
		}
	}

	if i.ProduceCode != "" {
		if valid, _ := regexp.Match(`^([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}$`, []byte(i.ProduceCode)); !valid {
			return errInvalidProductCode
		}
	}

	i.UnitPrice = math.Round(i.UnitPrice*100) / 100

	return nil
}
