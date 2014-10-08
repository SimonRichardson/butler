package butler

var (
	emptyString      = String("")
	constEmptyString = Constant(emptyString)
)

type String string

func NewString(value string, f func(Any) Any) String {
	chars := String(value).ToList()
	worth := chars.Head().Map(
		func(head Any) Any {
			return chars.FoldLeft(emptyString, func(acc, c Any) Any {
				return f(c).(String) + acc.(String)
			}).(String)
		},
	).GetOrElse(constEmptyString)
	return worth.(String)
}

func (s String) ToList() List {
	num := len(string(s))
	res := make([]Any, num, num)
	for i := 0; i < num; i++ {
		res[i] = String(s[i])
	}
	return SliceToList(res)
}
