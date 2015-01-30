package http

import (
	"encoding/json"
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

func (s alphaLowerString) Make(rand *rand.Rand, size int) string {
	var (
		rnd = rand.Intn(size) + 1
		buf = make([]rune, rnd)
		num = len(alphaLower)
	)
	for k, _ := range buf {
		buf[k] = alphaLower[rand.Intn(num)]
	}
	return string(buf)
}

func (s alphaLowerString) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(alphaLowerString(s.Make(rand, size)))
}

func (s alphaLowerString) String() string {
	return string(s)
}

type alphaString string

func (s alphaString) Make(rand *rand.Rand, size int) string {
	var (
		rnd = rand.Intn(size) + 1
		buf = make([]rune, rnd)
		num = len(alpha)
	)
	for k, _ := range buf {
		buf[k] = alpha[rand.Intn(num)]
	}
	return string(buf)
}

func (s alphaString) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(alphaString(s.Make(rand, size)))
}

func (s alphaString) String() string {
	return string(s)
}

type alphaNumericString string

func (s alphaNumericString) Make(rand *rand.Rand, size int) string {
	var (
		rnd = rand.Intn(size) + 1
		buf = make([]rune, rnd)
		num = len(alphaNumeric)
	)
	for k, _ := range buf {
		buf[k] = alphaNumeric[rand.Intn(num)]
	}
	return string(buf)
}

func (s alphaNumericString) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(alphaNumericString(s.Make(rand, size)))
}

func (s alphaNumericString) String() string {
	return string(s)
}

type name struct {
	Name string `json:"name"`
}

func (s name) Generate(rand *rand.Rand, size int) reflect.Value {
	x := name{}
	x.Name = alphaNumericString("").Make(rand, size)
	return reflect.ValueOf(x)
}

func (s name) String() string {
	x, _ := json.MarshalIndent(s, "", "\t")
	return string(x)
}

func fail(_ g.Any) g.Any {
	return fmt.Errorf("Fail")
}
