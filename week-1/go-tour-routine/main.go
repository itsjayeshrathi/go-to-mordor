package main

import (
	"fmt"
	"sync"
	"time"
)

func Say(word string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func Sum(s []int, c chan int) {
	sum := 0
	for _, ele := range s {
		sum += ele
	}
	c <- sum
}

func Add() {
	sum := 0
	for i := 1; i <= 10000; i++ {
		sum += i
	}
	fmt.Println("inside Addde", sum)
}
func main() {
	var sg sync.WaitGroup
	for range 1 {
		fmt.Println("inside main")
		go Add()
	}
	sg.Wait()
	fmt.Println("done")

}
