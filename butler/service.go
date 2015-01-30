package butler

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type service struct {
	request  list
	response list
	callable func() g.Any
}

func Service(request, response Builder, callable func() g.Any) service {
	return service{
		request:  newList(request.List()),
		response: newList(response.List()),
	}
}

func (s service) Compile() g.Either {
	fmt.Println(s.request.Build().Run())
	return g.Either_.Of(g.Empty{})
}

func (s service) String() string {
	return getRoute(s.request.list).Fold(
		func(x g.Any) g.Any {
			return fmt.Sprintf("Service(`%s`)", x)
		},
		g.Constant("Service()"),
	).(string)
}

type list struct {
	list g.List
}

func newList(x g.List) list {
	return list{
		list: x,
	}
}

func (r list) Build() g.WriterT {
	return g.WriterT_.Sequence(r.list.Map(func(x g.Any) g.Any {
		return AsBuild(x).Build()
	}))
}
