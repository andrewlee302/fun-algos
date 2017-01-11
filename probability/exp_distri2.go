package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// @Whose Algo to simulate the Exponential
// Distribution and Poisson Distribution
func main() {
	rand.Seed(time.Now().UnixNano())
	lambda := 1.0 / 40

	event_cnt := 0
	try_interval := 1.0 // try every second
	try_times := 10000
	for i := 0; i < try_times; i++ {
		if rand.Float64() < change(lambda, try_interval) {
			// shot
			event_cnt++
		}
	}

	// the mean should be approximate the 1/lambda
	fmt.Println(float64(try_times) * try_interval / float64(event_cnt))
}

// calculate the probability of (X <= rv) in exp distribution
// @param lambda, the param of the exp
// @param rv, the upper value
// @return probability of (X<=v)
func change(lambda float64, v float64) float64 {
	return 1.0 - math.Pow(math.E, -(lambda*v))
}
