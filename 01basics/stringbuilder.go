package basics

import (
	"log"
	"strings"
)

var src = []string{"Back", "To", "The", "Future", "Part", "III"}

// this is inefficient program, should consider using strinbs.Builder type instead
func StringConcat() {
	var title string
	for i, word := range src {
		if i != 0 {
			title += word
		}
	}
	log.Println(title)
}

// code sample of use of strings.Builder
func StringBuilder() {
	var builder strings.Builder
	builder.Grow(1000)
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
}
