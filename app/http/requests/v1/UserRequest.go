package v1

import (
	"net/http"
	"net/url"
	"github.com/samayamnag/boilerplate/app/rules"

	"github.com/thedevsaddam/govalidator"
)

func CreateUserRequest(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"full_name": []string{"required", "min:3", "max:50"},
		"email":     []string{"required", "min:8", "max:50", "email"},
		"password":  []string{"required", "between:3,20"},
	}

	messages := govalidator.MapData{
		"full_name": []string{"required:Full Name can not be blank"},
		"email":     []string{"required:Email can not be blank", "email:Invalid E-mail"},
		"password":  []string{"required:Password can not be blank", "between:Password length should between 3 and 20"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)

	return v.Validate()
}

func UpdateUserRequest(r *http.Request) url.Values {

	// Load custom validation rules
	rules.CustomValidationRules(r)

	rules := govalidator.MapData{
		"full_name": []string{"required", "min:3", "max:50"},
		"email":     []string{"required", "min:8", "max:50", "email", "unique:users"},
		"password":  []string{"required", "between:3,20"},
	}

	messages := govalidator.MapData{
		"full_name": []string{"required:Full Name can not be blank"},
		"email":     []string{"required:Email can not be blank", "email:Invalid E-mail"},
		"password":  []string{"required:Password can not be blank", "between:Password length should between 3 and 20"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)

	return v.Validate()
}
