package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
}
