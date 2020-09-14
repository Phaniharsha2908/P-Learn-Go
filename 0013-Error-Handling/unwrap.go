package main

import (
	"errors"
	"fmt"
)

var FeesNotSubmitted = errors.New("fees not submitted")
var AdmissionCancelled = errors.New("admission not possible")
var foo = errors.New("foo")

func fees() error {
	return fmt.Errorf("%w", FeesNotSubmitted)
}

func admission() error {
	return fmt.Errorf("%w %v ", fees(), AdmissionCancelled)
}

func foo1() error {
	return fmt.Errorf("%w %v", admission(), foo)
}

func main() {
	err := foo1()
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

}
