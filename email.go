package minimizers

import (
	fake "github.com/brianvoe/gofakeit"
)


func MinimizeEmail(level MinimizationLevel, email interface{}) interface{} {
	if level == MinimizationAnonymize {
		return fake.Email()
	}

	return email
}
