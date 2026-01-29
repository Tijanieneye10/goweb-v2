package validations

type Error map[string][]string

func (e Error) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e Error) Get(field string) []string {
	return e[field]
}
