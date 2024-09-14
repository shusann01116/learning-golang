package main

import (
	"fmt"
	"runtime/debug"
)

//nolint:forbidigo
func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	fmt.Println(info)
}
