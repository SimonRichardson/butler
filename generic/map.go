package generic

type Map interface {
	Map(func(Any) Any) Map
}
