package validations

import (
	"fmt"
	"net/url"
	"strings"
)

type error map[string][]string

func (e error) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e error) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

func (e error) Has(field string) bool {
	return len(e[field]) > 0
}

type Form struct {
	url.Values
	Error error
}

func NewForm(data url.Values) *Form {
	return &Form{
		Values: data,
		Error:  error(make(map[string][]string)),
	}
}

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Error.Add(field, fmt.Sprintf("%s is required", field))
		}
	}

	return f
}

func (f *Form) MaxLength(field string, n int) *Form {
	value := f.Get(field)
	if len(value) > n {
		f.Error.Add(field, fmt.Sprintf("%s can't be more than %v", field, n))
	}
}
