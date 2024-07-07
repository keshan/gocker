package main

import (
	"fmt"
	"log"
	"os"

	"github.com/keshan/gocker/pkg/container"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide a command")
	}

	switch os.Args[1] {
	case "run":
		run()
	default:
		log.Fatal("Unknown command")
	}
}

func run() {
	fmt.Println("Creating container...")
	container, err := container.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Container created successfully")
	fmt.Printf("Container ID: %s\n", container.ID)
}