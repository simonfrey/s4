package s4

import (
	"bytes"
	"github.com/hashicorp/vault/shamir"
	"github.com/pkg/errors"
	"github.com/simonfrey/s4/crypto"
)

var splitCode []byte = []byte("\n*=_=_=_=*\n\n")

// DistributeBytes takes the given in bytes and distributes them to n shares.
// At least k shares are required to restore the initial data. For better performance and security use DistributeBytesAES
func DistributeBytes(in []byte, n, k uint64) ([][]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}
	return shamir.Split(in, int(n), int(k))
}

// RecoverBytes recovers the given shares generate by DistributeBytes to their original payload.
func RecoverBytes(in [][]byte) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}
	return shamir.Combine(in)
}

// DistributeBytesAES takes the given in bytes and distributes them to n shares.
// At least k shares are required to restore the initial data
// In comparison to DistributeBytes this function uses AES on the payload and only distributes the key. This is a lot
// fast, as AES is highly optimized and backed by hardware support in most modern systems. The downside is a massive increase
// in share size, as every share now also has to contain the full AES payload.
func DistributeBytesAES(in []byte, n, k uint64) ([][]byte, error) {
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

	finalShares := make([][]byte, len(byteShares))
	for k, byteShare := range byteShares {
		finalShares[k] = append(append(byteShare, splitCode...), ciphertext...)
	}
	return finalShares, nil
}

// RecoverBytesAES recovers the given shares generate by DistributeBytesAES to their original payload.
func RecoverBytesAES(in [][]byte) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}

	var cipherText []byte
	shares := make([][]byte, len(in))
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
