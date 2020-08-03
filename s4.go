package s4

import (
	"bytes"
	"github.com/hashicorp/vault/shamir"
	"github.com/pkg/errors"
	"s4/crypto"
)

type Shares [][]byte

var splitCode []byte = []byte("\n*=_=_=_=*\n\n")

func DistributeBytes(in []byte, n, k uint64) (Shares, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}
	return shamir.Split(in, int(n), int(k))
}

func RecoverBytes(in Shares) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}
	return shamir.Combine(in)
}

func DistributeBytesAES(in []byte, n, k uint64) (Shares, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}

	key := crypto.NewEncryptionKey()

	byteShares, err := DistributeBytes(key[:], n, k)
	if err != nil {
		return nil, errors.Wrap(err, "could not distribute bytes")
	}

	ciphertext, err := crypto.Encrypt(in, key)
	if err != nil {
		return nil, errors.Wrap(err, "could not aes encrypt input")
	}

	finalShares := make(Shares, len(byteShares))
	for k, byteShare := range byteShares {
		finalShares[k] = append(append(byteShare, splitCode...), ciphertext...)
	}
	return finalShares, nil
}

func RecoverBytesAES(in Shares) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}

	var cipherText []byte
	shares := make(Shares, len(in))
	for k, s := range in {
		s := bytes.Split(s, splitCode)
		if len(s) != 2 {
			return nil, errors.New("invalid aes base")
		}

		if len(cipherText) == 0 {
			cipherText = s[1]
		} else if !bytes.Equal(cipherText, s[1]) {
			return nil, errors.New("AES cipher text differs between shares")
		}

		shares[k] = s[0]
	}

	recoveredKey, err := RecoverBytes(shares)
	if err != nil {
		return nil, errors.Wrap(err, "could not recover bytes of key")
	}
	if len(recoveredKey) != 32 {
		return nil, errors.New("recovered key is not size 32 byte")
	}

	key := [32]byte{}
	for k, v := range recoveredKey {
		key[k] = v
	}

	clearText, err := crypto.Decrypt(cipherText, &key)
	if err != nil {
		return nil, errors.Wrap(err, "could not aes decrypt input")
	}

	return clearText, nil
}
