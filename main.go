package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readMaxNumber(invalid bool) int {
	clearTerminal()

	if invalid {
		fmt.Println("Invalid input. The number have to be greater than 1")
	}

	var maxNum int
	fmt.Println("Enter the maximum number: ")
	fmt.Scan(&maxNum)

	if maxNum < 2 {
		return readMaxNumber(true)
	}

	return maxNum
}

func trialDivision(maxNum int) {
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

	fmt.Println(primes)
}

func main() {
	maxNum := readMaxNumber(false)

	trialDivision(maxNum)
}
