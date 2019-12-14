package minimizers

import "github.com/adaptant-labs/go-minimizer/datagen"

// Tokenize produces a random token for a given string-based data field, prefixed with 'tok_'.
func Tokenize() (string, error) {
	return datagen.GenerateRandomToken()
}
