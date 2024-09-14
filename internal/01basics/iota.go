//go:generate stringer -type=CarOption
//go:generate stringer -type=MyEnum

package basics

type MyEnum int

const (
	Apple MyEnum = iota + 1
	Banana
	PineApple
)

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)
