package bigint

import (
	"fmt"
	"math/big"
)

var maxSize *big.Int

func init() {
	maxSize = big.NewInt(0)
	maxSize.Exp(big.NewInt(2), big.NewInt(127), nil)
	maxSize.Sub(maxSize, big.NewInt(1))
}

func BytesToBigIntSlice(in []byte) ([]*big.Int, error) {
	base := big.NewInt(0)
	base.SetBytes(in)

	res := make([]*big.Int, 0)

	for base.Cmp(maxSize) >= 0 {
		modRes := big.NewInt(0)
		modRes.Mod(base, maxSize)

		base.Sub(base, modRes)
		base.Div(base, maxSize)

		res = append(res, modRes)
	}

	res = append(res, base)

	return res, nil
}

func BigIntSliceToBytes(in []*big.Int) ([]byte, error) {
	base := in[len(in)-1]

	for k := len(in) - 2; k >= 0; k-- {
		i := in[k]
		if i == nil {
			return nil, fmt.Errorf("Got nil pointer to big int")
		}

		base.Mul(base, maxSize)
		base.Add(base, i)
	}

	return base.Bytes(), nil
}
