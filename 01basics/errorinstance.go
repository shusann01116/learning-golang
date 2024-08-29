package basics

import "errors"

// これは厳密な実装
type errDataBase int

func (e errDataBase) Error() string {
	return "Database Error"
}

const (
	ErrDataBase errDataBase = 0
)

// var を使った気軽な定義
var ErrConnError = errors.New("basics: connection is already closed")
