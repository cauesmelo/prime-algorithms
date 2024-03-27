package main

import (
	"math"

	"github.com/cauesmelo/prime-algorithms/internal/benchmark"
)

func trialDivision(maxNum int) []int {
	primes := make([]int, 0)

	for i := 2; i < maxNum; i++ {
		isPrime := true
		sqrt := math.Sqrt(float64(i))

		for _, prime := range primes {
			if float64(prime) > sqrt {
				break
			}

			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func sieveOfEratosthenes(maxNum int) []int {
	arr := make([]bool, maxNum-1)
	const positionShift = 2

	for n, notPrime := range arr {
		if notPrime {
			continue
		}

		number := n + positionShift
		currMultiple := number + number

		for currMultiple <= maxNum {
			multiplePosition := currMultiple - positionShift

			arr[multiplePosition] = true

			currMultiple += number
		}
	}

	primes := make([]int, 0)

	for idx, notPrime := range arr {
		if notPrime {
			continue
		}

		primes = append(primes, idx+positionShift)
	}

	return primes
}

// TODO: Fix implementation
func dijkstra(maxNum int) []int {
	type PoolItem struct {
		n int
		m int
	}

	primes := []int{2}
	pool := []PoolItem{{n: 2, m: 4}}

	for i := 3; i < maxNum; i++ {
		hasMultiple := false

		for poolIdx := range pool {
			if pool[poolIdx].m == i {
				pool[poolIdx].m += pool[poolIdx].n

				hasMultiple = true
			}
		}

		if !hasMultiple {
			primes = append(primes, i)
			pool = append(pool, PoolItem{n: i, m: int(math.Pow(float64(i), 2))})
		}
	}

	return primes
}

func main() {
	maxNum := 90000

	b := &benchmark.Benchmark{}

	b.Start("Trial Division")
	trialDivision(maxNum)
	b.EndPrint()

	b.Start("Sieve of Eratosthenes")
	sieveOfEratosthenes(maxNum)
	b.EndPrint()

	b.Start("Dijkstra")
	dijkstra(maxNum)
	b.EndPrint()
}
