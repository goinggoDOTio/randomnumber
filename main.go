package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Since rand produces deterministic sequences, we need to seed it with different values to initialize the default source.
	//So I am using the current time in nano seconds since 1 of january of 1970
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 1; i <= 10; i++ {
		fmt.Println("Range:", i*10, "Random number:", rand.Intn(i*10))
	}
}
