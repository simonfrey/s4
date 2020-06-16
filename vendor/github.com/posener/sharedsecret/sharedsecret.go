// Package sharedsecret is implementation of Shamir's Secret Sharing algorithm.
//
// Shamir's Secret Sharing is an algorithm in cryptography created by Adi Shamir. It is a form of
// secret sharing, where a secret is divided into parts, giving each participant its own unique
// part. To reconstruct the original secret, a minimum number of parts is required. In the threshold
// scheme this number is less than the total number of parts. Otherwise all participants are needed
// to reconstruct the original secret.
// See (wiki page) https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing.
package sharedsecret

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/posener/sharedsecret/internal/polynom"
)

// prime128 is a large prime number that fits into 128 bits (value of 2^127 - 1).
var prime128 = prime128Value()

// Share is a part of a secret.
type Share struct {
	x, y *big.Int
}

// New creates n Shares and a secret. k defines the minimum number of shares that should be
// collected in order to recover the secret. Recovering the secret can be done by calling Recover
// with more than k Share objects.
func New(n, k int64) (shares []Share, secret *big.Int) {
	return distribute(nil, n, k)
}

// Distribute creates n Shares for a given secret. k defines the minimum number of shares that
// should be collected in order to recover the secret. Recovering the secret can be done by calling
// Recover with more than k Share objects.
func Distribute(secret *big.Int, n, k int64) (shares []Share) {
	shares, _ = distribute(secret, n, k)
	return shares
}

// distribute creates n shares. The secret argument is optional. It returns the shares and the
// secret for the shares.
func distribute(secret *big.Int, n, k int64) ([]Share, *big.Int) {
	if n < k {
		panic("irrecoverable: not enough shares to reconstruct the secret.")
	}
	if k <= 0 {
		panic("number of shares must be positive.")
	}
	p := polynom.NewRandom(k, prime128)

	// Set the first coefficient to the secret (the value at x=0) if the secret was given. And
	// anyway store the first coefficient in the secret variable.
	if secret != nil {
		if secret.Cmp(prime128) > 0 {
			panic("secret value is too big (must be lower than 2^127 - 1)")
		}
		p.SetCoeff(0, secret)
	}
	secret = p.Coeff(0)

	// Create the shares which are the value of p at any point but x != 0. Choose x in [1..n].
	shares := make([]Share, 0, n)
	for i := int64(1); i <= n; i++ {
		x := big.NewInt(i)
		y := p.ValueAt(x)
		shares = append(shares, Share{x: x, y: y})
	}

	return shares, secret
}

// Recover the secret from shares. Notice that the number of shares that is used should be at least
// the recover amount (k) that was used in order to create them in the New function.
func Recover(shares ...Share) (secret *big.Int) {
	// Convert the shares to a list of points x[i], y[i].
	xs := make([]*big.Int, len(shares))
	ys := make([]*big.Int, len(shares))
	for i := range shares {
		xs[i] = shares[i].x
		ys[i] = shares[i].y
	}
	// Evaluate the polynom that goes through all (x[i], y[i]) points at x=0.
	return polynom.Interpolate(big.NewInt(0), xs, ys, prime128)
}

// String dumps the share object to a string.
func (s Share) String() string {
	return s.x.String() + "," + s.y.String()
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s Share) MarshalText() ([]byte, error) {
	x, err := s.x.MarshalText()
	if err != nil {
		return nil, err
	}
	y, err := s.y.MarshalText()
	if err != nil {
		return nil, err
	}
	return append(append(x, ','), y...), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *Share) UnmarshalText(txt []byte) error {
	parts := bytes.Split(txt, []byte{','})
	if len(parts) != 2 {
		return errors.New("expected two parts")
	}
	s.x = &big.Int{}
	s.y = &big.Int{}
	err := s.x.UnmarshalText(parts[0])
	if err != nil {
		return err
	}
	return s.y.UnmarshalText(parts[1])
}

// prime128 returns a large prime that fits into 128 bits. It is 12th Mersenne Prime. (for this
// application we want a known prime number as close as possible to our security level; e.g. desired
// security level of 128 bits -- too large and all the ciphertext is large; too small and security
// is compromised) It is equal to 2^127 - 1. (13th Mersenne Prime is 2^521 - 1).
func prime128Value() *big.Int {
	p := big.NewInt(2)
	p.Exp(p, big.NewInt(127), nil)
	p.Sub(p, big.NewInt(1))
	return p
}
