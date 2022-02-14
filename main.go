package main

import (
	"fmt"
)

func makePanic() {
	defer func() {
		value := recover()
		err, ok := value.(error)
		if !ok {
			fmt.Printf("captured user panic: %s", value)
		}
		if err != nil {
			fmt.Printf("unexpected panic captured: %s", value)
		}

	}()

	a := 0
	q := 1 / a
	fmt.Printf("%v", q)

	panic("expected panic")

}

func main() {

	makePanic()

}
