package interfaces

import (
	"context"
	"log"
)

type TypeAssersionContext string

//nolint:ineffassign,wastedassign,gosmopolitan
func anyfunc() {
	var v any = "何でも入る変数"
	v = 1
	v = 3.14192

	var slices = []any{
		"関ヶ原",
		1600,
	}

	var ieyasu = map[string]any{
		"名前":  "徳川家康",
		"生まれ": 1543,
	}

	_ = []any{v, slices, ieyasu}
}

func typeAssertion() {
	ctx := context.WithValue(context.Background(), TypeAssersionContext("favorite"), "aiueo")

	// Type assertion
	// must use with ok to check if was the assertion succeed
	if s, ok := ctx.Value("favorite").(string); ok {
		log.Printf("My favorite thing is %s\n", s)
	}

	// Type switch
	switch v := ctx.Value("favorite").(type) {
	case string:
		log.Printf("string\n")
	case int:
		log.Printf("number\n")
	case complex128:
		log.Printf("complex\n")
	default:
		log.Printf("unknown number: %v\n", v)
	}
}
