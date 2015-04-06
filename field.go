package forms

import (
	"errors"
	"fmt"
)

type FormatterFunc func(string) string

type Field struct {
	Name       string
	formatters []FormatterFunc
	loader     LoaderFunc
	validators []ValidatorFunc
	required   bool
	empty      interface{}
}

func (f *Field) Validate(rawValue string) (interface{}, error) {
	if rawValue == "" {
		if f.required {
			return nil, errors.New(fmt.Sprintf("%s is required", f.Name))
		} else {
			return f.empty, nil
		}
	}

	// format raw input
	rawValue = f.format(rawValue)

	// deserialize
	value, err := f.loader(rawValue)
	if err != nil {
		return nil, err
	}

	// validate
	err = f.validate(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (f *Field) format(rawValue string) string {
	for _, f := range f.formatters {
		rawValue = f(rawValue)
	}
	return rawValue
}

func (f *Field) validate(value interface{}) error {
	for _, v := range f.validators {
		if err := v(value); err != nil {
			return err
		}
	}
	return nil
}
