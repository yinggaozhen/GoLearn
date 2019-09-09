package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @link https://colobu.com/2015/10/09/Linux-Signals/
func main() {
	signalChan := make(chan os.Signal, 1)
	signalOutputChan := make(chan string, 1)

	signal.Notify(signalChan)

	// signal recv
	go func() {
		for {
			if sig := <-signalChan; sig != nil {
				fmt.Println(sig)

				switch sig {
				case syscall.SIGINT:
					signalOutputChan <- "Interrupt"
					os.Exit(-1)
					return
				case syscall.SIGSTOP:
					signalOutputChan <- "loop stop"
				default:
					signalOutputChan <- "Default"
				}
			}
		}
	}()

	// signal send
	go func() {
		for {
			fmt.Println("send signal")
			signalChan <- syscall.SIGSTOP
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		if v, fetched := <- signalOutputChan; fetched {
			fmt.Println(v)
		}

		time.Sleep(time.Second)
	}
}
