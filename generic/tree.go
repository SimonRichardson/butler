package generic

type Tree interface {
	Chain(func(Any) Tree) Tree
	Map(func(Any) Any) Tree
	FoldLeft(Any, func(Any, Any) Any) Any
	Merge(t Tree) Tree
	Children() Option
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
				clean = b.Filter(func(x Any) bool {
					node = x.(TreeNode)
					return node.value != val
				})
				others = b.Filter(func(x Any) bool {
					node = x.(TreeNode)
					return node.value == val
				})
				children = AsList(others.FoldLeft(NewNil(), func(a, b Any) Any {
					return AsList(a).Concat(b.(TreeNode).nodes)
				}))
			)
			return List_.Of(
				NewTreeNode(
					Option_.Of(val),
					rec(nodes, children),
				),
			).Concat(clean)
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

func (t TreeNode) Children() Option {
	return Option_.Of(t.nodes)
}

type TreeNil struct {
}

func NewTreeNil() TreeNil {
	return TreeNil{}
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

func (t TreeNil) Children() Option {
	return Option_.Empty()
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
