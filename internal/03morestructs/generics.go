package morestructs

type Struct[T any] struct {
	T T
}

// 03morestructs/generics.go:9:26: syntax error: method must have no type parameters
// func (s Struct[T]) Method[R any](r R) {
// 	fmt.Println(s.t, r)
// }
