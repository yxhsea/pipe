package main

import (
	"fmt"
)

func main(){
	ch := make(chan string)
	for i := 0; i < 1000 ; i++ {
		go PrintHelloWorld(i,ch)
	}

	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

func PrintHelloWorld(i int, ch chan string)  {
	for {
		ch <- fmt.Sprintf("hello world %d\n",i)
	}
}
