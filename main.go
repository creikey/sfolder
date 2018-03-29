package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		client, err := CreateClient(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		test, _, err := client.RWriter.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(test))
	}
}
