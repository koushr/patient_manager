package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println("shutting down, caused by", sig)
		close(done)
	}()

	<-done
	fmt.Println("Graceful shutdown.")
}
