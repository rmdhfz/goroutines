package belajar_golang_goroutines

import (
	"testing"
	"time"
	"fmt"
)

func TestAtomic(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func(){
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}