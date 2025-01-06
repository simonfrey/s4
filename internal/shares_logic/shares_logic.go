package shares_logic

import (
	"encoding/base64"
	"fmt"
	"github.com/simonfrey/s4/pkg/base24"
	"github.com/simonfrey/s4/pkg/format"
	"github.com/simonfrey/s4/pkg/s4"
)

func RecoverShares(shares []string) (data []byte, err error) {
	useAES := false
	version := float32(-1)
	byteShares := make([][]byte, len(shares))

	for k, v := range shares {
		if len(v) == 0 {
			return nil, fmt.Errorf("please provide all shares. Share '%d' is empty", (k + 1))
		}

		f, err := format.ParseTravelFormat(v)
		if err != nil {
			return nil, fmt.Errorf("could not parse travel format: %w", err)
		}

		var tmpBytes []byte
		switch f.Version {
		case 0.6:
			tmpBytes, err = base24.DecodeStringToBytes(f.Data)
			if err != nil {
				return nil, fmt.Errorf("could not decode base24 string: %w", err)
			}
		case 0.5:
			tmpBytes, err = base64.StdEncoding.DecodeString(f.Data)
			if err != nil {
				return nil, fmt.Errorf("could not decode base64 string: %w", err)
			}
		default:
			return nil, fmt.Errorf("version not supported")
		}

		if f.UseAES {
			useAES = true
		}
		if version == -1 {
			version = f.Version
		}
		if f.Version != version {
			return nil, fmt.Errorf("version mismatch between shares")
		}
		byteShares[k] = tmpBytes
	}

	var clearText []byte
	if useAES {
		clearText, err = s4.RecoverBytesAES(byteShares)
	} else {
		clearText, err = s4.RecoverBytes(byteShares)
	}
	if err != nil {
		return nil, fmt.Errorf("could not recover bytes: %w", err)
	}

	return clearText, nil
}

func DistributeShares(inBytes []byte, n, k int, useAES, useBase24 bool) (shares []string, err error) {
	var byteShares [][]byte
	if useAES {
		byteShares, err = s4.DistributeBytesAES(inBytes, uint64(n), uint64(k))
	} else {
		byteShares, err = s4.DistributeBytes(inBytes, uint64(n), uint64(k))
	}
	if err != nil {
		return nil, fmt.Errorf("could not distribute bytes: %w", err)
	}

	shares = make([]string, len(byteShares))
	for k, byteShare := range byteShares {
		f := format.Format{
			UseAES:                    useAES,
			Version:                   0.5,
			OptimizedHumandReadbility: false,
		}
		if useBase24 {
			base24Encoded, err := base24.EncodeBytesToString(byteShare)
			if err != nil {
				return nil, fmt.Errorf("could not encode base24 string: %w", err)
			}
			f.Data = base24Encoded
			f.Version = 0.6
			f.OptimizedHumandReadbility = true
		} else {
			f.Data = base64.StdEncoding.EncodeToString(byteShare)
		}
		shares[k] = format.CreateTravelFormat(f)
	}

	return shares, nil
}
