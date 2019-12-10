package minimizers

import (
	fake "github.com/brianvoe/gofakeit"
	"strings"
)

// MinimizeName provides name minimization. Coarse-grained minimization returns
// first initial and last name, while fine-grained minimization returns only
// the first name. Anonymization will return a randomly generated name.
func MinimizeName(level MinimizationLevel, name interface{}) interface{} {
	if level == MinimizationAnonymize {
		return fake.Name()
	}

	if len(name.(string)) == 0 || level == MinimizationNone {
		return name
	}

	parts := strings.Split(name.(string), " ")

	switch level {
	case MinimizationCoarse:
		parts[0] = string([]rune(parts[0])[0])
		return strings.Join(parts, " ")
	case MinimizationFine:
		return parts[0]
	}

	return name
}
