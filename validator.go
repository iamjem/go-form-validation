package forms

import (
	"errors"
	"regexp"
)

type ValidatorFunc func(interface{}) error

func ReValidator(pattern string, errorMsg string) ValidatorFunc {
	validRe := regexp.MustCompile(pattern)
	return ValidatorFunc(func(value interface{}) error {
		strValue, ok := value.(string)
		// check type assertion
		if !ok {
			return errors.New("Invalid string value")
		}
		// check regular expression
		if !validRe.MatchString(strValue) {
			return errors.New(errorMsg)
		}
		return nil
	})
}

func RangeValidator(min, max int, errorMsg string) ValidatorFunc {
	return ValidatorFunc(func(value interface{}) error {
		intValue, ok := value.(int)
		// check type assertion
		if !ok {
			return errors.New("Invalid int value")
		}
		// check range
		if intValue < min || intValue > max {
			return errors.New(errorMsg)
		}
		return nil
	})
}
