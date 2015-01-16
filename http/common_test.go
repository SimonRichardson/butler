package http

import (
	"fmt"
	"math/rand"
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

var (
	alphaLower   = []rune("abcdefghijklmnopqrstuvwxyz")
	alpha        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type alphaLowerString string

func (s alphaLowerString) Generate(rand *rand.Rand, size int) reflect.Value {
	var (
		rnd = rand.Intn(50) + 1
		buf = make([]rune, rnd)
		num = len(alphaLower)
	)
	for k, _ := range buf {
		buf[k] = alphaLower[rand.Intn(num)]
	}
	return reflect.ValueOf(alphaLowerString(string(buf)))
}

func (s alphaLowerString) String() string {
	return string(s)
}

type alphaString string

func (s alphaString) Generate(rand *rand.Rand, size int) reflect.Value {
	var (
		rnd = rand.Intn(50) + 1
		buf = make([]rune, rnd)
		num = len(alpha)
	)
	for k, _ := range buf {
		buf[k] = alpha[rand.Intn(num)]
	}
	return reflect.ValueOf(alphaString(string(buf)))
}

func (s alphaString) String() string {
	return string(s)
}

type alphaNumericString string

func (s alphaNumericString) Generate(rand *rand.Rand, size int) reflect.Value {
	var (
		rnd = rand.Intn(50) + 1
		buf = make([]rune, rnd)
		num = len(alphaNumeric)
	)
	for k, _ := range buf {
		buf[k] = alphaNumeric[rand.Intn(num)]
	}
	return reflect.ValueOf(alphaNumericString(string(buf)))
}

func (s alphaNumericString) String() string {
	return string(s)
}

func fail(_ g.Any) g.Any {
	return fmt.Errorf("Fail")
}
