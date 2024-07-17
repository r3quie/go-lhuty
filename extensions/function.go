package binding

func NewFunctionBoundString(funcedString string, input String, input2 String) String {
	var str string = funcedString
	input.AddListener(&boundString{val: &str})
	input2.AddListener(&boundString{val: &str})
	return &boundString{val: &str}
}

/*
type twostringsToFunc struct {
	base

	format string

	from  String
	from2 String
}

type fn func(string, string) string

func TwoStringsToFunc(str String, str2 String, function fn) String {
	v := &twostringsToFunc{from: str, from2: str2}
	str.AddListener(v)
	str2.AddListener(v)

}
*/
