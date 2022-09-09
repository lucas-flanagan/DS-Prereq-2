package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Initialize slice with random values
func Generate_Randoms(x int) []int {
	slce := make([]int, x)
	for i := 0; i < x; i++ {
		slce[i] = rand.Intn(100)
	}
	return slce
}

// Calculate the sum normally, complexity O(n)
func NormalSum(slce []int) (int, time.Duration) {
	sum := 0
	start := time.Now()
	for _, v := range slce {
		sum += v
	}
	duration := time.Since(start)
	return sum, duration
}

// Gochannel driver functions, each complexity O(n)
func ParallelSum(c chan int, slce []int, wg *sync.WaitGroup) {
	defer wg.Done()
	localsum := 0
	for _, v := range slce {
		localsum += v
	}
	c <- localsum
}

// Compute the sum using two parallel go channels.
// Each channel will be responsible for computing the sum of half the slice's values
func Compute_Parallel(slce []int) (int, time.Duration) {
	var wg sync.WaitGroup
	wg.Add(2)
	sum := 0
	c := make(chan int, 2)
	start := time.Now()
	go ParallelSum(c, slce[:len(slce)/2], &wg)
	go ParallelSum(c, slce[len(slce)/2:], &wg)
	wg.Wait()
	close(c)
	for val := range c {
		sum += val
	}
	duration := time.Since(start)
	return sum, duration
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	slce := Generate_Randoms(x)
	normal_sum, normal_duration := NormalSum(slce)
	parallel_sum, parallel_duration := Compute_Parallel(slce)
	fmt.Println("Normal Sum: ", normal_sum)
	fmt.Println("Normal Sum Duration: ", normal_duration)
	fmt.Println("Parallel Sum: ", parallel_sum)
	fmt.Println("Parallel Sum Duration: ", parallel_duration)
}
