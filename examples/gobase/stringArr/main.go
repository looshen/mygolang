package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	stringArray()
	producerConsumer()
}

func stringArray() {
	arr := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(arr)
	for key, value := range arr {
		if strings.EqualFold(value, "stupid") {
			arr[key] = "smart"
		} else if arr[key] == "weak" {
			arr[key] = "strong"
		}
	}
	fmt.Println(arr)
}

func producerConsumer() {
	q := make(chan int, 10)
	go func(c chan int) {
		rand.Seed(time.Now().UnixNano())
		for {
			r := rand.Intn(100)
			fmt.Printf("producer data:%+v\n", r)
			c <- r
			time.Sleep(time.Second)
		}
	}(q)
	for {
		// time.Sleep(time.Second * 3)
		data := <-q
		fmt.Printf("consumer data:%+v\n", data)
	}
}
