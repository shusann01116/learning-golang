package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil)).Elem()

	fmt.Printf("os.File implements io.Writer: %v\n", fileType.Implements(writerType))
	// os.File implements io.Writer: false
}
