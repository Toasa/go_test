package main

import (
	   "fmt"
)

type Duck interface{
	 bark()
}

type Dog struct {
	 bark_str string
}

func (d Dog) bark() {
	 fmt.Println(d.bark_str)
}

func bark_test(d Duck) {
	 d.bark()
}

func main() {
	 dog := Dog{"bowbow"}
	 bark_test(dog)
}