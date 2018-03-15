package api

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
	"myracloud.com/myra-shell/api/vo"
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
