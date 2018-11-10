package rules

import (
	"net/http"
	"github.com/thedevsaddam/govalidator"
)

func CustomValidationRules(r *http.Request) {
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		return nil
	})

}