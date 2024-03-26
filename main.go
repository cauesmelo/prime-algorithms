package main

import (
	"fmt"
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

func main() {
	maxNum := readMaxNumber(false)

	fmt.Println(maxNum)
}
