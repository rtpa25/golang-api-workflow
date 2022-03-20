package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/rtpa25/go_api_worflow/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//check the validity of the currency
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
