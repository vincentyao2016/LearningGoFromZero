package test

import (
	"fmt"
	"testing"
	"time"
)

// https://medium.com/better-programming/understanding-golang-and-goroutines-72ac3c9a014d
func TestGoroutionWithMultiCores(t *testing.T) {
	// scale up on multiple cores
	// runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution without sleep: " + elapsedTime.String())

	time.Sleep(10 * time.Second)

	elapsedTime = time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
}
