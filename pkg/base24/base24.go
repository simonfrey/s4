package base24

import (
	"bytes"
	"fmt"
	"gitlab.com/phil9909/base24"
	"strings"
)

func EncodeBytesToString(inBytes []byte) (string, error) {
	// Append empty bytes to make the length of inBytes a multiple of 4
	if overflow := len(inBytes) % 4; overflow != 0 {
		for i := 0; i < 4-overflow; i++ {
			inBytes = append(inBytes, 0)
		}
	}

	enc, err := base24.StdEncoding.EncodeToString(inBytes)
	if err != nil {
		return "", fmt.Errorf("could not encode bytes to base24: %w", err)
	}
	return strings.ToLower(enc), nil
}

func DecodeStringToBytes(inString string) ([]byte, error) {
	outBytes, err := base24.StdEncoding.DecodeString(strings.TrimSpace(inString))
	if err != nil {
		return nil, err
	}

	return bytes.TrimRight(outBytes, string([]byte{0})), nil
}
