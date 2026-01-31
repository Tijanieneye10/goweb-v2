package validations

import (
	"fmt"
	"net/mail"
	"net/url"
	"strings"
	"unicode/utf8"
)

type customError map[string][]string

func (e customError) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e customError) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

func (e customError) Has(field string) bool {
	return len(e[field]) > 0
}

type Form struct {
	url.Values
	Error customError
}

func NewForm(data url.Values) *Form {
	return &Form{
		Values: data,
		Error:  customError(make(map[string][]string)),
	}
}

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Add(field, fmt.Sprintf("%s is required", field))
		}
	}

	return f
}

func (f *Form) MaxLength(field string, n int) *Form {
	value := f.Get(field)
	if utf8.RuneCountInString(value) > n {
		f.Error.Add(field, fmt.Sprintf("%s can't be more than %d", field, n))
	}

	return f
}

func (f *Form) MinLength(field string, n int) *Form {
	value := f.Get(field)
	if utf8.RuneCountInString(value) < n {
		f.Error.Add(field, fmt.Sprintf("%s can't be more than %v", field, n))
	}

	return f
}

func (f *Form) Email(field string) *Form {
	value := f.Get(field)

	if value == "" {
		f.Error.Add(field, "Email address is required")
	}

	_, err := mail.ParseAddress(value)

	if err != nil {
		f.Error.Add(field, "must be a valid email address")
	}

	return f
}

func (f *Form) Valid() bool {
	return len(f.Error) == 0
}
