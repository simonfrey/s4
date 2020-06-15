package s4

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	ss "github.com/posener/sharedsecret"
	"log"
	"math/big"
	"s4/bigint"
)

type Shares [][]byte
type Rows [][]ss.Share

func DistributeBytes(in []byte, n, k uint64) (Shares, error) {
	bIntInputSlice, err := bigint.BytesToBigIntSlice(in)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := distributeBigIntSlice(bIntInputSlice, n, k)
	if err != nil {
		log.Fatal(err)
	}

	return rowsToShares(rows), nil
}

func RecoverBytes(in Shares) ([]byte, error) {
	if len(in) == 0 {
		return nil, errors.New("in must not be empty")
	}

	bIntSlice, err := recoverToBigIntSlice(sharesToRows(in))
	if err != nil {
		return nil, errors.Wrap(err, "could not recover big int slice")
	}
	allNil := true
	for _, v := range bIntSlice {
		if v != nil {
			allNil = false
			break
		}
	}
	if allNil {
		return nil, errors.New("could not recover big int slice: all recovered shares are nill")
	}

	return bigint.BigIntSliceToBytes(bIntSlice)
}

func distributeBigIntSlice(in []*big.Int, n, k uint64) (Rows, error) {
	res := make([][]ss.Share, len(in))

	for ik, v := range in {
		res[ik] = ss.Distribute(v, int64(n), int64(k))
	}
	return res, nil
}

func recoverToBigIntSlice(in Rows) ([]*big.Int, error) {
	res := make([]*big.Int, len(in))

	for ik, v := range in {
		res[ik] = ss.Recover(v...)
	}
	return res, nil
}

func revertShareSlices(in [][]ss.Share) [][]ss.Share {
	out := make([][]ss.Share, len(in[0]))
	for _, row := range in {
		for k, col := range row {
			out[k] = append(out[k], col)
		}
	}
	return out
}

func bytesToShares(in [][]byte) [][]ss.Share {
	out := make([][]ss.Share, len(in))

	for k, b := range in {
		scanner := bufio.NewScanner(bytes.NewReader(b))
		for scanner.Scan() {
			tmp := ss.Share{}
			err := tmp.UnmarshalText(scanner.Bytes())
			if err != nil {
				panic(err)
			}
			out[k] = append(out[k], tmp)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return out

}
func sharesToBytes(in [][]ss.Share) [][]byte {
	out := make([][]byte, len(in))
	for k, share := range in {
		for _, shareRow := range share {
			out[k] = append(out[k], []byte(shareRow.String()+string('\n'))...)
		}
	}
	return out
}
func sharesToRows(in Shares) Rows {
	if len(in) == 0 {
		return Rows{}
	}
	return revertShareSlices(bytesToShares(in))
}

func rowsToShares(in Rows) Shares {
	if len(in) == 0 {
		return Shares{}
	}
	return sharesToBytes(revertShareSlices(in))
}
