package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP)

	go func() {
		for {
			s := <-sigs

			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CRTL-C")
				fmt.Println("Some clean clean up is occuring")
				cleanUp()
				done <- struct{}{}

			case syscall.SIGSTOP:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CRTL-Z")
				fmt.Println("Some clean clean up is occuring")
				cleanUp()
				done <- struct{}{}
			}
		}
	}()

	fmt.Println("Program is blocked until a signal is caught(crtl-z, crtl-c)")
	done <- struct{}{}
	fmt.Println("Out of here")
}

func cleanUp() {
	fmt.Println("Simulating cleanup")

	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting Files... Not really", i)
		time.Sleep(1 * time.Second)
	}
}
