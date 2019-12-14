package minimizers

import (
	"github.com/adaptant-labs/go-minimizer/datagen"
	"unicode"
)

// Mask randomly swaps out alphanumeric values within a string with type-matching random values. The format (and case)
// of the string is preserved, while special characters are left in-place.
// e.g. Mask("123-456-7890") => "155-503-1463"
func Mask(input string) (string, error) {
	output := []rune(input)
	for i, c := range input {
		if unicode.IsLetter(c) {
			output[i] = datagen.GenerateRandomLetter(unicode.IsUpper(c))
		} else if unicode.IsNumber(c) {
			output[i] = datagen.GenerateRandomDigit()
		}
	}

	return string(output), nil
}

// MaskWithPattern provides a masking variant in which all alphanumeric characters are replaced with a given pattern.
// e.g. MaskWithPattern("123-456-7890", 'X') => "XXX-XXX-XXXX".
func MaskWithPattern(input string, pattern rune) (string, error) {
	output := []rune(input)
	for i, c := range input {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			output[i] = pattern
		}
	}

	return string(output), nil
}
