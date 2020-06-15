package bigint

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestBigSliceConversion(t *testing.T) {
	rand.Seed(time.Now().Unix())

	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	bIntSlice, err := BytesToBigIntSlice(input)
	if err != nil {
		log.Fatal(err)
	}

	nBytes, err := BigIntSliceToBytes(bIntSlice)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(input, nBytes) {
		log.Fatal("Input and output bytes differ")
	}
}
