package main

import (
	"fmt"
	"runtime"
	"time"
)

func basicGoroutine() {
	runtime.GOMAXPROCS(4) //выделение ядер процессора
	start := time.Now()
	func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)

	start = time.Now()
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime = time.Since(start)

	defer fmt.Println("Total Time For Execution: " + elapsedTime.String())
	defer time.Sleep(time.Second)
}
