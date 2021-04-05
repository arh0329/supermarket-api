package models

import (
	"errors"
	"math"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	errInvalidProductCode = errors.New("ProductCode is invalid must be only alphanumeric characters. The produce codes are sixteen characters long, with dashes separating each four character group")
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Item struct {
	Name        string  `json:"name" validate:"required,alphanum"`
	ProduceCode string  `json:"produceCode" validate:"required"`
	UnitPrice   float64 `json:"unitPrice" validate:"required,numeric"`
}

func (i *Item) Validate() (err error) {
	err = validate.Struct(i)
	if err != nil {
		return err
	}

	if i.ProduceCode != "" {
		if valid, _ := regexp.Match(`^([a-zA-Z0-9]{4}-){3}[a-zA-Z0-9]{4}$`, []byte(i.ProduceCode)); !valid {
			return errInvalidProductCode
		}
	}

	i.UnitPrice = math.Round(i.UnitPrice*100) / 100

	return nil
}
