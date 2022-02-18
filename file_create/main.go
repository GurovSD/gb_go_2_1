package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	for i := 0; i < 10; i++ {
		rn := rand.Intn(10)
		file, err := os.Create("/home/laooglee/gb/go/file_" + strconv.Itoa(rn))
		if err != nil {
			fmt.Printf("file creation failed at attempt %v with value %v", i, rn)
			os.Exit(1)
		}

		defer func() {
			file.Write([]byte(fmt.Sprintf("file creation failed at attempt %v with value %v", i, rn)))
			file.Close()
		}()

		res := 10 / rn
		file.Write([]byte(strconv.Itoa(res)))

		fmt.Println(rn)
	}

}
