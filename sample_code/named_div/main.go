package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainderConfusing(numerator, denominator int) (result int, remainder int,
	err error) {
	// assign some values
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)

	result, remainder, err := divAndRemainderConfusing(5, 2)
	fmt.Println(result, remainder, err)
}
