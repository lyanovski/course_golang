package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var destination string
	// Won't work on the Playground since the time is frozen.
	rand.Seed(time.Now().Unix())
	name := []string{
		"Fahmi",
		"Lyanovski",
		"Meigi",
		"Rothdam",
	}
	n := rand.Int() % len(name)
	fmt.Print("Enter your destination: ")
	fmt.Scanf("%s", &destination)
	fmt.Print("you will meet ", name[n], " in ", destination)
}
