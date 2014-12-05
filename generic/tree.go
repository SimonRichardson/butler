package generic

type Tree interface {
	Children() Option
	Chain(func(Any) Tree) Tree
	Map(func(Any) Any) Tree
	FoldLeft(Any, func(Any, Any) Any) Any
	Merge(t Tree) Tree
	Match(func(List, Any, int) bool) List
}

type TreeNode struct {
	value Any
	nodes List
}

func NewTreeNode(x Any, y List) TreeNode {
	return TreeNode{
		value: x,
		nodes: y,
	}
}

func (t TreeNode) Children() Option {
	return Option_.Of(t.nodes)
}

func (t TreeNode) Chain(f func(Any) Tree) Tree {
	x := f(t.value)
	if _, ok := x.(TreeNil); ok {
		return x
	}

	return NewTreeNode(x, t.nodes.Chain(func(x Any) List {
		return List_.Of(x.(Tree).Chain(f))
	}))
}

func (t TreeNode) Map(f func(Any) Any) Tree {
	return t.Chain(func(x Any) Tree {
		return Tree_.Of(f(x))
	})
}

func (t TreeNode) FoldLeft(x Any, f func(Any, Any) Any) Any {
	return t.nodes.FoldLeft(f(x, t.value), func(x, y Any) Any {
		return y.(Tree).FoldLeft(x, f)
	})
}

func (t TreeNode) Merge(m Tree) Tree {
	var rec func(a, b List) List
	rec = func(a, b List) List {
		return a.Chain(func(x Any) List {
			if _, ok := x.(TreeNil); ok {
				return List_.Of(x)
			}

			var (
				node  = x.(TreeNode)
				val   = node.value
				nodes = node.nodes
				tuple = b.Partition(func(x Any) bool {
					if _, ok := x.(TreeNil); ok {
						return false
					}

					node = x.(TreeNode)
					return node.value == val
				})
				fst      = AsList(tuple.Fst())
				children = AsList(fst.FoldLeft(List_.Empty(), func(a, b Any) Any {
					return AsList(a).Concat(b.(TreeNode).nodes)
				}))
			)
			return List_.Of(
				NewTreeNode(
					Option_.Of(val),
					rec(nodes, children),
				),
			).Concat(AsList(tuple.Snd()))
		})
	}
	return NewTreeNode(
		Option_.Empty(),
		rec(
			List_.Of(t),
			List_.Of(m),
		),
	)
}

func (t TreeNode) Match(f func(List, Any, int) bool) List {
	var rec func(a List, b Tree, c int) List
	rec = func(a List, b Tree, c int) List {
		if _, ok := b.(TreeNil); ok {
			return a
		}

		var (
			x = b.(TreeNode)
			y = Option_.FromBool(f(a, x.value, c), x.value)
		)

		return AsList(y.Fold(
			func(y Any) Any {
				return x.nodes.FoldLeft(NewCons(x.value, a), func(a, b Any) Any {
					if _, ok := b.(TreeNil); ok {
						return a
					}
					var (
						list = AsList(a)
						node = b.(TreeNode)
					)
					return rec(list, node, c+1)
				})
			},
			Constant(a),
		))
	}
	return rec(NewNil(), t, 0)
}

type TreeNil struct {
}

func NewTreeNil() TreeNil {
	return TreeNil{}
}

func (t TreeNil) Children() Option {
	return Option_.Empty()
}

func (t TreeNil) Chain(f func(Any) Tree) Tree {
	return t
}

func (t TreeNil) Map(f func(Any) Any) Tree {
	return t
}

func (t TreeNil) FoldLeft(x Any, f func(Any, Any) Any) Any {
	return x
}

func (t TreeNil) Merge(m Tree) Tree {
	return m
}

func (t TreeNil) Match(f func(List, Any, int) bool) List {
	return List_.Empty()
}

// Static methods

var (
	Tree_ = tree{}
)

type tree struct{}

func (x tree) Of(v Any) Tree {
	return NewTreeNode(v, NewNil())
}

func (x tree) Empty() Tree {
	return NewTreeNil()
}

func (x tree) FromList(l List) Tree {
	return AsTree(l.Reverse().FoldLeft(NewTreeNil(), func(a, b Any) Any {
		node := AsTree(a)
		return NewTreeNode(b, List_.Of(node))
	}))
}
