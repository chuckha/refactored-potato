package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	prisoners = 100
	// Set this to -1 for a random seed.
	seed = int64(1525554535)
)

func main() {
	rand.Seed(getSeed())
	n := prisoners
	labels := r(n)
	printSlice(labels)
	names := make([]int, n)
	copy(names, labels)
	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})
	printSlice(names)
	draw(labels, names)
}

func getSeed() int64 {
	if seed != -1 {
		return seed
	}
	x := time.Now().Unix()
	fmt.Println("Seed: ", x)
	return x
}

// tell the prisoners to draw.
func draw(labels, values []int) {
	for _, prisoner := range r(prisoners) {
		current := prisoner
		count := 0
		done := false
		for !done {
			fmt.Printf("%v→", current)
			current = values[current-1]
			count++
			if current == prisoner {
				fmt.Printf("%v\n%v found their number in box %v after %v picks\n\n", prisoner, prisoner, current, count)
				done = true
				continue
			}
			if count >= prisoners/2 {
				fmt.Printf("everyone dies after %v picks\n.", count)
				os.Exit(0)
			}
		}
	}
}

// return a slice of int from 1 to n
func r(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i + 1
	}
	return out
}

// mostly ignore from here down.

// all numbers must be >= 0
func maxWidth(input []int) int {
	max := -1
	for _, n := range input {
		if n > max {
			max = n
		}
	}
	base := 1
	l := 0
	for {
		if max/base == 0 {
			return l
		}
		l++
		base *= 10
	}
}

func printSlice(input []int) {
	maxw := maxWidth(input)
	format := fmt.Sprintf("%%%dv ", maxw)
	for _, v := range input {
		fmt.Printf(format, v)
	}
	fmt.Println()
}
