package http

import (
	"fmt"
	"strings"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type Route struct {
	doc.Api
	path String
}

func NewRoute(path string) Route {
	return Route{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected route `%s`"),
			doc.NewInlineText("Unexpected route `%s`"),
		)),
		path: NewString(path, PathChar()),
	}
}

func (r Route) Build() g.WriterT {
	var (
		extract = func(a g.Any) g.WriterT {
			var (
				x = g.AsTuple3(a)
				y = AsString(x.Fst())
			)
			return g.WriterT_.Of(y.String()).
				Tell(fmt.Sprintf("Extract `%v`", y))
		}
		compile = func(a g.Any) g.WriterT {
			x := r.Route().Bimap(g.Constant1(a), g.Constant1(a))
			return g.WriterT_.Lift(x).
				Tell(fmt.Sprintf("Compile `%v`", a))
		}
		api = func(x doc.Api) func(g.Either) g.WriterT {
			return func(y g.Either) g.WriterT {
				return g.WriterT_.Lift(x.Run(y)).
					Tell(fmt.Sprintf("Api `%v`", y))
			}
		}
		finalize = func(a Route) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(name g.WriterT) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					split = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							return strings.Split(b.(string), "/")
						}
					}
					unshift = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							return b.([]string)[1:]
						}
					}
					withConditions = func(parts []string) func(g.List, g.Any, int) bool {
						return func(a g.List, b g.Any, c int) bool {
							return g.AsOption(b).Fold(
								func(a g.Any) g.Any {
									if c < len(parts) {
										return AsPathNode(a).Match(parts[c])
									}
									return false
								},
								g.Constant(c == 0),
							).(bool)
						}
					}
					match = func(x g.Any) func(g.Any) g.Any {
						return func(y g.Any) g.Any {
							return r.Route().Bimap(
								func(x g.Any) g.Any {
									return g.List_.Empty()
								},
								func(x g.Any) g.Any {
									var (
										a = g.AsTree(x)
										b = withConditions(y.([]string))
									)
									return g.NewTuple2(y, g.Walker_.Match(a, b))
								},
							)
						}
					}
					correct = func(x g.Any) func(g.Any) g.Any {
						return func(y g.Any) g.Any {
							var (
								put = func(z g.Any) g.Any {
									return g.NewTuple2(a, a)
								}
								a = g.AsEither(y)
							)
							return g.AsEither(a.Fold(
								g.Either_.Left,
								func(x g.Any) g.Any {
									var (
										a = g.AsTuple2(x)
										b = len(a.Fst().([]string))
										c = g.AsList(a.Snd())
									)
									return g.Either_.FromBool(b == c.Size(), c)
								},
							)).Bimap(put, put)
						}
					}
					flatten = func(a g.Any) g.StateT {
						return g.NewStateT(g.AsEither(a))
					}
					program = g.StateT_.Of(a).
						Chain(modify(split)).
						Chain(g.Get()).
						Chain(modify(unshift)).
						Chain(g.Get()).
						Chain(modify(match)).
						Chain(g.Get()).
						Chain(modify(correct)).
						Chain(g.Get()).
						Chain(flatten)
				)
				return g.AsTuple2(a).Append(program)
			}
		}

		path = r.path.Build()

		program = path.
			Chain(extract).
			Chain(compile)
	)

	return join(program, api(r.Api), func(x g.Any) []g.Any {
		return singleton(x)
	}).Bimap(
		finalize(r),
		finalize(r),
	).Bimap(
		matcher(path),
		matcher(path),
	)
}

func (r Route) Route() g.Either {
	return compilePath(r.String()).Bimap(
		func(x g.Any) g.Any {
			return g.Tree_.Empty()
		},
		func(x g.Any) g.Any {
			return g.Tree_.FromList(g.AsList(x))
		},
	)
}

func (r Route) String() string {
	return r.path.String()
}
