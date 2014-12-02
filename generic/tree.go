package generic

type Tree interface {
	Chain(func(Any) Tree) Tree
	Map(func(Any) Any) Tree
	FoldLeft(Any, func(Any, Any) Any) Any

	Children() Option
}

type TreeNode struct {
	Value Any
	nodes List
}

func NewTreeNode(x Any, y List) TreeNode {
	return TreeNode{
		Value: x,
		nodes: y,
	}
}

func (t TreeNode) Chain(f func(Any) Tree) Tree {
	x := f(t.Value)
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
	return t.nodes.FoldLeft(f(x, t.Value), func(x, y Any) Any {
		return y.(Tree).FoldLeft(x, f)
	})
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
