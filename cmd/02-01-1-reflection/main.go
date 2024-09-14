package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"reflect"
)

func main() {
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil)).Elem()

	slog.Info(fmt.Sprintf("os.File implements io.Writer: %v\n", fileType.Implements(writerType)))
	// os.File implements io.Writer: false
}
