// Package base24 implements base24 encoding as specified by https://www.kuon.ch/post/2020-02-27-base24/
package base24

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

// Encoding Most of the time you want to use StdEncoding, but you can provide a custom alphabet using NewEncoding
type Encoding struct {
	encode    [24]byte
	decodeMap [256]byte
}

const encodeStd = "ZAC2B3EF4GH5TK67P8RS9WXY"

// NewEncoding returns an Encoding defined by the given alphabet,
// The alphabet must be 24 distinct characters
func NewEncoding(encoder string) *Encoding {
	if len(encoder) != 24 {
		panic("encoding alphabet is not 64-bytes long")
	}
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoding alphabet contains newline character")
		}
	}

	e := new(Encoding)
	copy(e.encode[:], encoder)

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}
	lowerCaseEncoder := strings.ToLower(encoder)
	upperCaseEncoder := strings.ToUpper(encoder)
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[lowerCaseEncoder[i]] = byte(i)
		e.decodeMap[upperCaseEncoder[i]] = byte(i)
	}
	return e
}

// StdEncoding is a base24 encoding with the alphabet defined in https://www.kuon.ch/post/2020-02-27-base24/
var StdEncoding = NewEncoding(encodeStd)

// Encode encodes src into dst.
// EncodedLen(len(src)) bytes will be written to dst.
// len(src) must be a multiple of 4.
// See https://www.kuon.ch/post/2020-02-27-base24/ for details
func (enc *Encoding) Encode(dst, src []byte) error {
	if len(src) == 0 {
		return nil
	}

	if len(src)%4 != 0 {
		return errors.New("length base24 data must be multiple of 4. See https://www.kuon.ch/post/2020-02-27-base24/")
	}
	// enc is a pointer receiver, so the use of enc.encode within the hot
	// loop below means a nil check at every operation. Lift that nil check
	// outside of the loop to speed up the encoder.
	_ = enc.encode

	for i := 0; i < len(src); i += 4 {
		chunk := binary.BigEndian.Uint32(src[i : i+4])
		//chunk := binary.LittleEndian.Uint32(src[i : i+4])
		for j := 0; j < 7; j++ {
			//dst[len(dst)-1-(i/4*7+j)] = enc.encode[chunk%24]
			dst[(i/4*7)+(6-j)] = enc.encode[chunk%24]
			chunk /= 24
		}
	}

	return nil
}

// EncodeToString returns the base24 encoding of src as string
func (enc *Encoding) EncodeToString(src []byte) (string, error) {
	buf := make([]byte, enc.EncodedLen(len(src)))
	err := enc.Encode(buf, src)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// EncodedLen returns the length in bytes of the base24 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(n int) int {
	return n / 4 * 7
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base24-encoded data.
func (enc *Encoding) DecodedLen(n int) int {
	return n / 7 * 4
}

// DecodeString returns the bytes represented by the base24 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
	dbuf := make([]byte, enc.DecodedLen(len(s)))
	n, err := enc.Decode(dbuf, []byte(s))
	return dbuf[:n], err
}

// Decode decodes src into dst.
// DecodedLen(len(src)) bytes will be written to dst.
// len(src) must be a multiple of 7.
// Might return an error when invalid data is given
// See https://www.kuon.ch/post/2020-02-27-base24/ for details
func (enc *Encoding) Decode(dst, src []byte) (n int, err error) {
	if len(src) == 0 {
		return 0, nil
	}

	if len(src)%7 != 0 {
		return 0, errors.New("length base24 encoded data must be multiple of 7. See https://www.kuon.ch/post/2020-02-27-base24/")
	}

	// Lift the nil check outside of the loop. enc.decodeMap is directly
	// used later in this function, to let the compiler know that the
	// receiver can't be nil.
	_ = enc.decodeMap

	dst[0] = enc.decodeMap[src[0]]

	k := 0
	for i := 0; i < len(src); i += 7 {
		chunk := uint32(0)
		for j := 0; j < 7; j++ {
			chunk *= 24
			decoded := enc.decodeMap[src[i+j]]
			if decoded == 255 {
				return k, fmt.Errorf(`illegal char '%s'`, string(src[i+j]))
			}
			chunk += uint32(decoded)
		}
		binary.BigEndian.PutUint32(dst[k:k+4], chunk)

		k += 4
	}

	return k, nil
}
