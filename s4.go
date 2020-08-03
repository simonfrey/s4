package s4

import (
	"github.com/hashicorp/vault/shamir"
	"github.com/pkg/errors"
)

type Shares [][]byte

var splitCode []byte = []byte("\n*=_=_=_=*\n\n")

func DistributeBytes(in []byte, n, k uint64) (Shares, error) {
	return shamir.Split(in, int(n), int(k))
}

func RecoverBytes(in Shares) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}
	return shamir.Combine(in)
}
