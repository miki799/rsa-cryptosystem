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

	n := new(big.Int).Mul(p, q)
	v := new(big.Int).Mul(new(big.Int).Sub(p, utils.ONE), new(big.Int).Sub(q, utils.ONE))

	k := selectK(v)

	d := new(big.Int).ModInverse(k, v)

	return &PublicKey{k, n}, &PrivateKey{d, n}, nil
}

func selectK(v *big.Int) *big.Int {
	k := big.NewInt(3)

	// Check if k and v are relatively prime gcd(k, v) = 1
	for new(big.Int).GCD(nil, nil, k, v).Cmp(big.NewInt(1)) != 0 {
		// increment k by 2 and check again
		k.Add(k, big.NewInt(2))
	}

	fmt.Println("Selected k:", k)

	return k
}

func Encrypt(message string, publicKey *PublicKey) (*big.Int, error) {

	c := new(big.Int).Exp(m, publicKey.k, publicKey.n)

	return c, nil
}
