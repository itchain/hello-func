package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Println("Error to read from stdin")
	}

	fmt.Print(string(bytes[:]) + " version : 4.0")
}
