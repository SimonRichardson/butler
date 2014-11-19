package output

import (
	"math/rand"
	"reflect"
	"testing/quick"
	"time"

	g "github.com/SimonRichardson/butler/generic"
)

func toEither(a []byte, b error) g.Either {
	if b != nil {
		return g.NewLeft(b)
	}
	return g.NewRight(a)
}

func generate(e Encoder) func(g.Any) g.Either {
	var (
		now = time.Now().Unix()
		rnd = rand.New(rand.NewSource(now))
	)
	return func(x g.Any) g.Either {
		var (
			typ     = reflect.TypeOf(x)
			val, ok = quick.Value(typ, rnd)
		)
		return g.Either_.FromBool(ok, val.Interface()).
			Chain(func(x g.Any) g.Either {
			return e.Encode(x)
		}).Map(func(x g.Any) g.Any {
			a := string(x.([]uint8))
			return a
		})
	}
}
