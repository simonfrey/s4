package base24_test

import (
	"fmt"
	"github.com/simonfrey/s4/pkg/base24"
	"testing"
)

func TestEncoding(t *testing.T) {
	inString := "Hello World"
	enc, err := base24.EncodeBytesToString([]byte(inString))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Encoded:", enc)

	dec, err := base24.DecodeStringToBytes(enc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Decoded:", string(dec))

	if inString != string(dec) {
		t.Errorf("Expected '%s' got '%s'", inString, dec)
	}
}
