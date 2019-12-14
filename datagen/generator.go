package datagen

import (
	"math/rand"
	"strconv"
	"strings"
)

const alphanumcharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	for i, c := range b {
		b[i] = alphanumcharset[c % byte(len(alphanumcharset))]
	}

	return string(b), nil
}

func GenerateRandomToken() (string, error) {
	token, err := GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	return "tok_" + token, nil
}

func GenerateRandomDigit() rune {
	n := strconv.Itoa(rand.Intn(10))
	return rune(n[0])
}

func GenerateRandomLetter(uppercase bool) rune {
	c := 'a' + rand.Intn(26)
	if uppercase {
		return rune(strings.ToUpper(string(c))[0])
	}

	return rune(c)
}
