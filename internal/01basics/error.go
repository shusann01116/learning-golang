package basics

import (
	"bufio"
	"errors"
	"os"

	"golang.org/x/exp/slog"
)

// これは上書きされてしまうが、厳格にエラーのIFを実装しようとするとヘビーなため
// プログラマの性善説を信じ var で定義している
// これは外部から `basics.ErrEOF = <new stmt>` で上書きできる
var ErrEOF = errors.New("EOF")

// Go でのエラーは値
type ErrorSample interface {
	Error() string
}

func ErrorHandling() {
	// don't use panic
	// panic()
	f, err := os.Open("important.txt")
	if err != nil {
		slog.Info("failed to open file")
		return
	}

	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		slog.Info("failed to read line")
		return
	}
	slog.Info(l)
}
