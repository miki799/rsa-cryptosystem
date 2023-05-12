package main

import (
	"fmt"

	"github.com/miki799/rsa-cryptosystem/rsa"
)

func main() {
	message, err := utils.ReadTextFromFile("message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("message to encrypt: %v\n", message)

	publicKey, privateKey, err := rsa.GenerateKeys(768)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("publicKey.String(): %v\n", publicKey.String())
	fmt.Printf("privateKey.String(): %v\n", privateKey.String())

	cryptogram, err := rsa.Encrypt(message, publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("cryptogram: %v\n", cryptogram)

}
