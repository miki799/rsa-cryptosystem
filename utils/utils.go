package utils

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

var ONE *big.Int = big.NewInt(1)
var TWO *big.Int = big.NewInt(2)

func ReadTextFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}

func GeneratePrimeNumber(bits int) (*big.Int, error) {
	for {
		num, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, fmt.Errorf("rand.Prime() returned error: %v", err)
		}

		if num.ProbablyPrime(bits) {
			return num, nil
		}

		fmt.Println("Generated number is not prime!", err)
	}
}
