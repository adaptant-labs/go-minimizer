package minimizers

import "strings"

type Errors []error

func (e Errors) Errors() []error {
	return e
}

func (e Errors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, ";")
}