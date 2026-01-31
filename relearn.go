package main

import (
	"fmt"
	"net/mail"
	"net/url"
	"unicode/utf8"
)

type customError map[string][]string

func (e customError) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e customError) Get(field string) string {
	if val, ok := e[field]; ok {
		return val[0]
	}

	return ""
}

type Form struct {
	url.Values
	Error customError
}

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		if utf8.RuneCountInString(field) > 1 {
			f.Add(field, fmt.Sprintf("%s is required", field))
		}
	}

	return f
}

func (f *Form) MinLength(field string, n int) *Form {
	value := f.Get(field)
	valueLen := utf8.RuneCountInString(value)

	if valueLen < n {
		f.Add(field, fmt.Sprintf("%s is required", field))
	}
	return f
}

func (f *Form) MaxLength(field string, n int) *Form {
	value := f.Get(field)
	valueLen := utf8.RuneCountInString(value)
	if valueLen > n {
		f.Add(field, fmt.Sprintf("%s is required", field))
	}

	return f
}

func (f *Form) Email(field string) *Form {
	value := f.Get(field)
	_, err := mail.ParseAddress(value)

	if err != nil {
		f.Add(field, fmt.Sprintf("%s is invalid", field))
	}

	return f

}
