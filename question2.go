package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// Initialize a slice with random integers
func Gen_Randoms(x int) []int {

	// Initialize slice and populate with random integers from 0-100
	// Source: gobyexample.com random-numbers
	slce := make([]int, x)
	for i := 0; i < x; i++ {
		slce[i] = rand.Intn(100)
	}
	return slce
}

// Define interface
type byValue []int

// Len, Less, and Swap functions are required definitions for Sort() and Stable()
func (a byValue) Len() int {
	return len(a)
}

func (a byValue) Less(i, j int) bool {
	return i < j
}

func (a byValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Sort the slice using go's func Sort(data Interface)
func SortSlice(slce []int) time.Duration {
	start := time.Now()
	sort.Sort(byValue(slce))
	sortDurr := time.Since(start)

	return sortDurr
}

// Sort the slice using go's func Stable(data Interface)
func StableSlice(slce []int) time.Duration {
	start := time.Now()
	sort.Stable(byValue(slce))
	stableDurr := time.Since(start)
	return stableDurr
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	slce := Gen_Randoms(x)
	SortSlice := SortSlice(slce)
	StableSlice := StableSlice(slce)
	fmt.Println("Sort duration for Slice: ", SortSlice)
	fmt.Println("Stable duration for Slice: ", StableSlice)
}
