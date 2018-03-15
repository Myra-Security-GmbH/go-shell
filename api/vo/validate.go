package vo

import validator "gopkg.in/go-playground/validator.v9"

//
// Validatable ...
//
type Validatable interface {
	Validate(validator.StructLevel)
}
