package validations

type Error map[string][]string

func (e Error) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e Error) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

func (e Error) Has(field string) bool {
	return len(e[field]) > 0
}