package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChanel := make(chan bool)
	go readUserInput(doneChanel)

	<-doneChanel
	close(doneChanel)
	fmt.Println("GoodBye")

}

func readUserInput(doneChanel chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		result, done := checkNumber(scanner)
		if done {
			doneChanel <- true
			return
		}
		fmt.Print(result)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numberToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print("Enter a whole number")
	}
	_, msg := isPrime(numberToCheck)

	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}

func intro() {
	fmt.Println("Is is prime?")
	fmt.Println("------------------>")
	fmt.Println("Enter a whole number And I will tell is it prime or Not!")
	prompt()
}

func prompt() {
	fmt.Print("----->")
}
