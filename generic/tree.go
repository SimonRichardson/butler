package generic

type Tree interface {
	Children() Option
	Chain(func(Any) Tree) Tree
	Map(func(Any) Any) Tree
	FoldLeft(Any, func(Any, Any) Any) Any
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

	y := x.(TreeNode)
	return NewTreeNode(y.value, t.nodes.Chain(func(x Any) List {
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

func (x tree) ToList(t Tree) List {
	return AsList(t.FoldLeft(NewNil(), func(a, b Any) Any {
		return NewCons(b, AsList(a))
	}))
}

func (x tree) Walker() walker {
	return Walker_
}

var (
	Walker_ = walker{}
)

type walker struct{}

func (w walker) Match(a Tree, f func(List, Any, int) bool) List {
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
	return rec(NewNil(), a, 0)
}

func (w walker) Map(a Tree, f func(Any, int, bool) Any) Tree {
	var rec func(a List, b int) List
	rec = func(a List, b int) List {
		return a.Chain(func(x Any) List {
			if _, ok := x.(TreeNil); ok {
				return List_.Of(x)
			}

			var (
				node     = x.(TreeNode)
				val      = node.value
				nodes    = node.nodes
				children = nodes.(Cons)
				last     = false
			)

			if nodes.Size() == 1 {
				_, last = children.head.(TreeNil)
			}

			return List_.Of(
				NewTreeNode(
					f(val, b, last),
					rec(nodes, b+1),
				),
			)
		})
	}
	return AsTree(rec(List_.Of(a), 0).Head().GetOrElse(func() Any {
		return NewTreeNil()
	}))
}

func (w walker) Merge(a, b Tree, f func(Any, Any) bool) Tree {
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
					return f(node.value, val)
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

	do := List_.Of(b)
	if _, ok := a.(TreeNil); ok {
		do = rec(
			List_.Of(a),
			List_.Of(b),
		)
	}

	return NewTreeNode(
		Option_.Empty(),
		do,
	)
}

func (w walker) Combine(a, b Tree, f func(List, List) Option) Tree {
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
					return f(node.value, val)
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

	do := List_.Of(b)
	if _, ok := a.(TreeNil); ok {
		do = rec(
			List_.Of(a),
			List_.Of(b),
		)
	}

	return NewTreeNode(
		Option_.Empty(),
		do,
	)
}
