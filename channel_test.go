package belajar_golang_goroutines

import (
	"testing"
	"fmt"
	"time"
	"strconv"
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

func OnlyIn(channel chan <- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hafiz Ramadhan"
}

func OnlyOut(channel <- chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Hafiz"
	channel <- "Ramadhan"

	fmt.Println("Selesai")
	fmt.Println(len(channel))
	fmt.Println(cap(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data: ",data)
	}

	fmt.Println("Selesai")
}