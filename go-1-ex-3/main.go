package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) 
	var eyes = rand.Intn(6) + 1      // Zahl ZWISCHEN 1 und6
	var when = time.Now()

	
	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)

	
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	fmt.Fprintln(eyesFile, eyes) // Only writing the number of eyes

	
	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating dice.log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintf(logFile, "Dice roll: %d eyes, rolled at: %s\n", eyes, when)
}
