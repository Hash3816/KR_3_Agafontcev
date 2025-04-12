package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	N := 5
	sum := 0

	set := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		value := rand.Intn(50) + 1
		set[value] = true
	}
	for el := range set {
		sum += el
	}

	P := float64(sum) / float64(N)
	for el := range set {
		if float64(el) > P {
			set[el] = false
		}
	}
	for el := range set {
		if set[el] == true {
			fmt.Println(el)
		}
	}
}
