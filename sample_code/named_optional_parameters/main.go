package main

import (
	"errors"
	"fmt"
	"os"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	if opts.Age < 0 {
		return errors.New("invalid age")
	}
	fmt.Print("Hello")
	if opts.FirstName != "" {
		fmt.Print(" ", opts.FirstName)
	}
	if opts.LastName != "" {
		fmt.Print(" ", opts.LastName)
	}
	if opts.Age != 0 {
		fmt.Print(", you are ", opts.Age, " years old")
	}
	fmt.Println(".")
	return nil
}

func main() {
	err := MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	err = MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
