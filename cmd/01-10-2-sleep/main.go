package main

import (
	"fmt"
	"time"
)

//nolint:forbidigo
func main() {
	fmt.Println("Start 3 sec sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("Start 3 sec end")

	fmt.Println("10 sec halt start")
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	<-timer.C
	fmt.Println("10 sec halt end")
}
