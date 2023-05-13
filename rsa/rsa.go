package rsa

import (
	"fmt"
	"math/big"

	"github.com/miki799/rsa-cryptosystem/utils"
)

type PublicKey struct {
	k *big.Int
	n *big.Int
}

func (pk PublicKey) String() string {
	return fmt.Sprintf("PublicKey {\n k: %v\n n: %v\n}", pk.k, pk.n)
}

type PrivateKey struct {
	d *big.Int
	n *big.Int
}

func (pk PrivateKey) String() string {
	return fmt.Sprintf("PrivateKey {\n d: %v\n n: %v\n}", pk.d, pk.n)
}

func GenerateKeys(bits int) (*PublicKey, *PrivateKey, error) {
	p, err1 := utils.GeneratePrimeNumber(bits)
	if err1 != nil {
		return nil, nil, err1
	}

	q, err2 := utils.GeneratePrimeNumber(bits)
	if err2 != nil {
		return nil, nil, err2
	}

	// n = p * q
	n := new(big.Int).Mul(p, q)

	// v = (p - 1)(q - 1)
	v := new(big.Int).Mul(new(big.Int).Sub(p, utils.ONE), new(big.Int).Sub(q, utils.ONE))

	// select small odd integer k such that gcd(k, v) = 1
	k := selectK(v)

	// select d such that (d * k) % v = 1
	d := new(big.Int).ModInverse(k, v)

	return &PublicKey{k, n}, &PrivateKey{d, n}, nil
}

func selectK(v *big.Int) *big.Int {
	k := big.NewInt(3)

	// Check if k and v are relatively prime gcd(k, v) = 1
	for new(big.Int).GCD(nil, nil, k, v).Cmp(utils.ONE) != 0 {
		// increment k by 2 and check again
		k.Add(k, utils.TWO)
	}

	return k
}

func Encrypt(message string, publicKey *PublicKey) []*big.Int {
	messageBigInts := utils.ConvertStringToBigIntsSlice(message)

	cryptogram := make([]*big.Int, 0)

	for i := 0; i < len(messageBigInts); i++ {
		c := new(big.Int)
		// (m[i] ^ k)modn
		c.Exp(messageBigInts[i], publicKey.k, publicKey.n)
		cryptogram = append(cryptogram, c)
	}

	return cryptogram
}

func Decrypt(cryptogram []*big.Int, privateKey *PrivateKey) string {
	encryptedMessage := make([]*big.Int, 0)

	for i := 0; i < len(cryptogram); i++ {
		m := new(big.Int)
		// (m[i] ^ d)modn
		m.Exp(cryptogram[i], privateKey.d, privateKey.n)
		encryptedMessage = append(encryptedMessage, m)
	}

	return utils.ConvertBigIntsSliceToString(encryptedMessage)
}

func Verify(originalMessage, encryptedMessage string) bool {
	return originalMessage == encryptedMessage
}
