package validations

var errors map[string][]string

func Add(field, message string) {
	errors[field] = append(errors[field], message)
}

func Get(field string) []string {
	return errors[field]
}
