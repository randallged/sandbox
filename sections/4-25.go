package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello everyone.")
	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("In foo!!")
}

// control flow
// (1) sequence
// (2) loop; interative
// (3) conditional
