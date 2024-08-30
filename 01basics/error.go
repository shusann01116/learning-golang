package basics

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// これは上書きされてしまうが、厳格にエラーのIFを実装しようとするとヘビーなため
// プログラマの性善説を信じ var で定義している
// これは外部から `basics.EOF = <new stmt>` で上書きできる
var EOF = errors.New("EOF")

// Go でのエラーは値
type errorSample interface {
	Error() string
}

func errorHandling() {
	// don't use panic
	// panic()
	f, err := os.Open("important.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}

	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read line")
		return
	}
	fmt.Println(l)
}
