package morestructs

import (
	"math/big"
	"sync"
)

type MyStruct struct {
	Value *big.Int
}

func Example2() {
	// declare `sync.Pool` struct
	pool := &sync.Pool{
		New: func() any {
			return &MyStruct{Value: nil}
		},
	}

	// retrieve the struct with type conversion
	// by Get() method
	str, ok := pool.Get().(*MyStruct)
	if !ok {
		panic("failed")
	}

	// put the struct back to pool for another uses
	pool.Put(str)
}
