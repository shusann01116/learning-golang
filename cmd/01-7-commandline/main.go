package main

import (
	"flag"
	"log"
)

var (
	FlagStr = flag.String("string", "default", "string flag")
	FlagInt = flag.Int("int", -1, "num flag")
)

func main() {
	flag.Parse()
	log.Println(*FlagStr)
	log.Println(*FlagInt)
	log.Println(flag.Args())
}
