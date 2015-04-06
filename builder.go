package forms

type FieldBuilder struct {
	formatters []FormatterFunc
	loader     LoaderFunc
	validators []ValidatorFunc
	required   bool
	empty      interface{}
}

func (fb *FieldBuilder) WithFormatters(formatters ...FormatterFunc) *FieldBuilder {
	fb.formatters = append(fb.formatters, formatters...)
	return fb
}

func (fb *FieldBuilder) WithValidators(validators ...ValidatorFunc) *FieldBuilder {
	fb.validators = append(fb.validators, validators...)
	return fb
}

func (fb *FieldBuilder) Loader(loader LoaderFunc) *FieldBuilder {
	fb.loader = loader
	return fb
}

func (fb *FieldBuilder) Required() *FieldBuilder {
	fb.required = true
	return fb
}

func (fb *FieldBuilder) Empty(value interface{}) *FieldBuilder {
	fb.empty = value
	return fb
}

func (fb *FieldBuilder) Build() *Field {
	return &Field{
		formatters: fb.formatters,
		loader:     fb.loader,
		validators: fb.validators,
		required:   fb.required,
		empty:      fb.empty,
	}
}
