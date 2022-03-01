package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filename := "./slice.go"
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file : %s. error message: %s\n", filename, err.Error())
	}
	reader := bufio.NewReader(fd)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatalf("can't read file : %s. error message: %s\n", filename, err.Error())
		}
		log.Println(string(line))
	}

}
