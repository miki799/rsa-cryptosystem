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

func ConvertStringToBigIntsSlice(str string) []*big.Int {
	strBytes := []byte(str)

	bigInts := make([]*big.Int, 0, len(strBytes))
	for _, b := range strBytes {
		bigInts = append(bigInts, big.NewInt(int64(b)))
	}

	return bigInts
}

func ConvertBigIntsSliceToString(arr []*big.Int) string {
	var result []byte

	for _, num := range arr {
		result = append(result, num.Bytes()...)
	}

	return string(result)
}
