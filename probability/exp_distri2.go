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
	try_interval := 1.0 // try every single second
	try_times := 10000
	for i := 0; i < try_times; i++ {
		if rand.Float64() < chance(lambda, try_interval) {
			// shot
			event_cnt++
		}
	}

	// the number of arrivals in 0-sum should be approximate the mean (sum * lambda)
	fmt.Println(float64(try_times)*try_interval*lambda, float64(event_cnt))
}

// calculate the probability of (X <= rv) in exp distribution.
// Because the memoryless property of exp distribution,
// Pr(X > nt | X > (n-1)t) = Pr(X > t).
// @param lambda, the param of the exp
// @param rv, the upper value
// @return probability of (X<=v)
func chance(lambda float64, v float64) float64 {
	return 1.0 - math.Pow(math.E, -(lambda*v))
}
