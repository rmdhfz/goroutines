package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"runtime"
)

func TestGetGomaxprocs(t *testing.T){

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}