package morestructs

import (
	"fmt"
	"time"
)

func example() {
	wait := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		wait <- struct{}{}
	}()

	fmt.Println("wait for 3 sec in go routine")
	<-wait
	fmt.Println("3 sec elappsed")
}
