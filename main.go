package main

import (
	"fmt"

	"time"

	"github.com/pkg/errors"
)

func divide(x int, y int) (res int, err error) {
	defer func() {
		value := recover()
		err, ok := value.(error)
		if !ok {
			err = errors.New(fmt.Sprintf("captured user panic: %s", value))
			err = errors.Wrapf(err, time.Now().Format("2006-01-02 15:04:02"))

			res = 0
			return
		}
		if err != nil {
			err = errors.Wrapf(err, "unexpected panic captured")
			err = errors.Wrapf(err, time.Now().Format("2006-01-02 15:04:02"))

			res = 0

			return
		}

	}()

	// res = float64(x) / float64(y)
	res = x / y

	// panic("expected panic")

	return res, nil

}

func main() {

	div, err := divide(2, 0)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(div)

}
