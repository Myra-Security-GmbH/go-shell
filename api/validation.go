package api

import (
	"fmt"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
	validator "gopkg.in/go-playground/validator.v9"
)

//
// validateVO ...
//
func validateVO(sl validator.StructLevel) {
	entity := sl.Current().Interface()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Given type does not implement vo.Validateable")
		}
	}()

	v := entity.(vo.Validatable)

	if v != nil {
		v.Validate(sl)
	}
}
