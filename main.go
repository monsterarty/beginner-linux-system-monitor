package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program starting...")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: %v", err)
	}
	fmt.Printf("Hostname: %v", hostname) 
}
