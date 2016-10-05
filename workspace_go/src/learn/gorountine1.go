package main

import (
	"fmt"
	"time"
)

var (
	Sig    = make(chan byte)
	strSig = make(chan string, 3)
)

func main() {
	// elements := []string{"A", "B", "C", "D"}
	// for _, element := range elements {
	// 	go func(e string) {
	// 		fmt.Println("ELement: " + e)
	// 	}(element)
	// }
	// time.Sleep(1 * time.Second)

	go gorountine1()
	go gorountine2()
	<-Sig
}

func gorountine1() {
	index := 0
	for {
		if index == 6 {
			close(strSig)
			fmt.Println("Sender: The Chan is closed")
			break
		}
		strSig <- "A"
		// time.Sleep(1 * time.Second)
		index += 1
		fmt.Printf("Sender: index-%d\n", index)
	}
}

func gorountine2() {
	for {
		v, ok := <-strSig
		if ok {
			fmt.Printf("recive: %#v\n", v)
			time.Sleep(5 * time.Second)
		} else {
			fmt.Printf("Reciver: will return\n")
			Sig <- '0'
		}
	}
}
