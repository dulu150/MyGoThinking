package main

import "fmt"

func MyRoutine(ch chan int) {
	fmt.Println("I'm do something,wait a second")
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go MyRoutine(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
