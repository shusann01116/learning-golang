package morestructs

type NoCopyStruct struct {
	self  *NoCopyStruct
	Value *string
}

func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		self:  nil,
		Value: &value,
	}
	r.self = r
	return r
}

func (n *NoCopyStruct) String() string {
	if n != n.self {
		panic("should not copy NoCopyStruct instance without Copy() method")
	}
	return *n.Value
}

func (n *NoCopyStruct) Copy() *NoCopyStruct {
	str := *n.Value
	p2 := &NoCopyStruct{
		self:  nil,
		Value: &str,
	}
	p2.self = p2
	return p2
}
