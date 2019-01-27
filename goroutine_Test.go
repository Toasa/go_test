package main

import (
	"fmt"
	"time"
)

func repeat(str string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str, "count: ", i)
	}
}

func main() {
	go repeat("a  ", 2)
	go repeat("b  ", 3)
	repeat("c  ", 6)

}
