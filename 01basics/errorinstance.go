package basics

import "errors"

// これは厳密な実装
type DataBaseError int

func (e DataBaseError) Error() string {
	return "Database Error"
}

const (
	ErrDataBase DataBaseError = 0
)

// var を使った気軽な定義
var ErrConnError = errors.New("basics: connection is already closed")
