package main

import (
	"runtime"
	"time"
)

func main() {
	subpDur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(subpDur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(subpDur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
