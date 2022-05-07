package belajar_golang_goroutines

import (
	"testing"
	"time"
	"fmt"
)

func TestTicker(t *testing.T){
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}