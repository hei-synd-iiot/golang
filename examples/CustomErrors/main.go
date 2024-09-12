package main

import (
	"errors"
	"fmt"
)

type argError struct { // (1)!
	arg     int
	message string
}

func (e *argError) Error() string { // (2)!
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"} // (3)!
	}
	return arg + 3, nil
}

func main() {

	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) { // (4)!
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
