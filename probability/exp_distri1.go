package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Donald Knuth algo to simulate the Exponential
// Distribution and Poisson Distribution
func main() {
	rand.Seed(time.Now().UnixNano())
	lambda := 1.0 / 40
	sum := 0.0
	times := 10000
	for i := 0; i < times; i++ {
		sum += nextTime(lambda)
	}
	// the number of arrivals in 0-sum should be approximate the mean (sum * lambda)
	fmt.Println(sum*lambda, float64(times))
}

// Interval between events satisfy the exponential distribution
// @param lambda indicate the event rate (times/sec)
// @return the interval (sec) till the next event
func nextTime(lambda float64) float64 {
	return -(math.Log(1.0 - rand.Float64())) / lambda
}
