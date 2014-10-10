package butler

var (
	emptyString      = String{Nil{}}
	constEmptyString = Constant(emptyString)
)

type String struct {
	List
}

func NewString(value string) String {
	chars := FromStringToList(value, func(s string) Any {
		return Char(s)
	})
	worth := chars.Head().Map(Constant1(chars)).GetOrElse(constEmptyString)
	return String{worth.(List)}
}

func (s String) String() string {
	return s.FoldLeft("", func(a, b Any) Any {
		return b.(Char).String() + a.(string)
	}).(string)
}

// Note : this should be a rune.
type Char string

func NewChar(value string) Char {
	return Char(value)
}

func (c Char) String() string {
	return string(c)
}
