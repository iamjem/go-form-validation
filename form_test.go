package forms

import (
	"errors"
	"net/url"
	"testing"
	"time"
)

func testForm() *Form {
	form := New()

	// username
	form.WithField("Username", new(FieldBuilder).
		Required().
		Loader(StringLoader).
		WithValidators(ReValidator(`^\w{5,}$`, "Username must be at least 3 characters.")))

	// password
	form.WithField("Password", new(FieldBuilder).
		Required().
		Loader(StringLoader).
		WithValidators(ReValidator(`^\w{10,}$`, "Password must be at least 10 characters.")))

	// dob
	form.WithField("DOB", new(FieldBuilder).
		Required().
		Loader(TimeLoader).
		WithValidators(ValidatorFunc(func(value interface{}) error {
		// require user is 18 or older
		if !value.(time.Time).Before(time.Now().AddDate(-18, 0, 1)) {
			return errors.New("Must be 18 or older to register")
		}
		return nil
	})))

	return form
}

func TestValidForm(t *testing.T) {
	formValues := url.Values{}
	formValues.Add("Username", "Johndoe")
	formValues.Add("Password", "SuperSecret")
	formValues.Add("DOB", time.Now().AddDate(-18, 0, 0).Format(time.RFC3339))

	form := testForm()
	valid := form.Valid(formValues)

	if !valid {
		t.Error("Form should be valid")
	}
}

func TestInvalidForm(t *testing.T) {
	formValues := url.Values{}
	formValues.Add("Username", "Foo")
	formValues.Add("Password", "TooShort")
	formValues.Add("DOB", time.Now().AddDate(-17, 0, 0).Format(time.RFC3339))

	form := testForm()
	valid := form.Valid(formValues)

	if valid {
		t.Error("Form should be invalid")
	}

	if form.Values != nil {
		t.Error("Form should not have values")
	}

	if form.Errors == nil {
		t.Error("Form should have errors")
	}

	for _, fieldName := range []string{"Username", "Password", "DOB"} {
		if _, ok := form.Errors[fieldName]; !ok {
			t.Errorf("Form should have errors for field %s", fieldName)
		}
	}
}
