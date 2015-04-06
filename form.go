package forms

import (
	"net/url"
)

type FormValues map[string]interface{}
type FormValidator func(formValues *FormValues) bool

type Form struct {
	Fields     map[string]*Field
	FieldNames []string
	Errors     map[string]error
	Values     FormValues
	validator  FormValidator
}

func New() *Form {
	return &Form{
		Fields:     make(map[string]*Field),
		FieldNames: make([]string, 0),
	}
}

// WithField adds the Field produced by the FieldBuilder to the Form under the given name.
func (f *Form) WithField(name string, fb *FieldBuilder) *Form {
	field := fb.Build()
	f.Fields[name] = field
	f.FieldNames = append(f.FieldNames, name)
	return f
}

// WithValidator adds the FormValidator to the form.
func (f *Form) WithValidator(validator FormValidator) *Form {
	f.validator = validator
	return f
}

// Valid validates each field followed by the form's validator if provided.
func (f *Form) Valid(postForm url.Values) bool {
	valid := true

	f.Errors = nil
	f.Values = nil

	formValues := make(FormValues)
	formErrors := make(map[string]error)

	// validate fields
	for _, fname := range f.FieldNames {
		fieldValue, fieldError := f.Fields[fname].Validate(postForm.Get(fname))
		if fieldError != nil {
			valid = false
			formErrors[fname] = fieldError
		} else {
			formValues[fname] = fieldValue
		}
	}

	// validate form
	if valid && f.validator != nil {
		valid = f.validator(&formValues)
	}

	// if its valid, make the values available
	// otherwise make the errors available
	if valid {
		f.Values = formValues
	} else {
		f.Errors = formErrors
	}

	return valid
}
