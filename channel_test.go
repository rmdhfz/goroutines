package belajar_golang_goroutines

import (
	"testing"
	"fmt"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Hafiz Ramadhan"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	data := <- channel
	fmt.Println(data)
	
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hafiz Ramadhan"
}

func TestAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	
	data := <- channel
	fmt.Println(data)
	
	time.Sleep(5 * time.Second)
}