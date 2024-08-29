package basics

type ErrorCode int

// `const` はコンパイル時に決定されるもののみ定義可能で、
// それを満たさないものはすべてコンパイルエラーとなる
const (
	f    int       = 10
	code ErrorCode = 10
)
