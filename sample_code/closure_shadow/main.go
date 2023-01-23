package main

import "fmt"

func main() {
	a := 20
	f := func() {
		fmt.Println(a)
		a := 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}
