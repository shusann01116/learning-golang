package morestructs

import (
	"log/slog"
	"time"
)

func Example() {
	wait := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		wait <- struct{}{}
	}()

	slog.Info("wait for 3 sec in go routine")
	<-wait
	slog.Info("3 sec elappsed")
}
