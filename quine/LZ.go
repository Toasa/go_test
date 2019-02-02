package main

import "fmt"

var input string;
var pos int;

func literal(n int) {
	for i := 0; i < n; i++ {
    	fmt.Printf("%c", input[pos]);
		pos++;
	}
}

func repeat(d, n int) {
	if pos - d < 0 {
	    fmt.Println("too large bytes to backward reading");
		return;
	}
	for i := pos - d; i < n; i++ {
		fmt.Printf("%c", input[i]);
	}
}

func main() {
    input = "this was a triumph";
	pos = 0;
	literal(7);
	repeat(4, 10);
	fmt.Println();
}