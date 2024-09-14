package basics

type ErrorCode int

// `const` はコンパイル時に決定されるもののみ定義可能で、
// それを満たさないものはすべてコンパイルエラーとなる
const (
	F    int       = 10
	Code ErrorCode = 10
)
