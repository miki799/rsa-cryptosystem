package main

import (
	"fmt"

	"github.com/miki799/rsa-cryptosystem/rsa"
	"github.com/miki799/rsa-cryptosystem/utils"
)

func main() {
	message, err := utils.ReadTextFromFile("message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("message to encrypt: %v\n", message)

	fmt.Println("Generating keys...")
	publicKey, privateKey, err := rsa.GenerateKeys(1024)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Encrypting...")
	cryptogram := rsa.Encrypt(message, publicKey)

	fmt.Println("Decrypting...")
	decryptedMessage := rsa.Decrypt(cryptogram, privateKey)

	if rsa.Verify(message, decryptedMessage) {
		fmt.Println("RSA cryptosystem works! Messages are the same!")
	} else {
		fmt.Println("RSA cryptosystem does not work! Messages are notthe same!")
	}
	fmt.Printf("decryptedMessage: %v\n", decryptedMessage)

}
