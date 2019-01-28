package main

import (
	"fmt"
	"time"
)

func main() {
	//	go count("sheep") //  1 processor
	//	go count("fish")  //  2 processors
	go count("dog") //  3 processors
	count("goat")   //  4 processors
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
