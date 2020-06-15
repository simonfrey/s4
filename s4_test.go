package s4

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestSecretSharingEqual(t *testing.T) {
	rand.Seed(time.Now().Unix())

	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n := uint64(rand.Intn(10-2) + 2)
	k := uint64(rand.Intn(int(n)-2) + 2)

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		log.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytes(restoreSlice)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(input, nBytes) {
		log.Fatal("Input and output bytes differ")
	}

}

func TestSecretSharingSameShares(t *testing.T) {
	rand.Seed(time.Now().Unix())

	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n := uint64(rand.Intn(10-2) + 2)
	k := uint64(rand.Intn(int(n)-2) + 2)

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		log.Fatal(err)
	}

	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		restoreSlice[i] = shares[0]
	}

	nBytes, err := RecoverBytes(restoreSlice)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(input, nBytes) {
		log.Fatal("Input and output bytes differ")
	}

}

func TestSecretSharingDistributeEmptyInput(t *testing.T) {
	rand.Seed(time.Now().Unix())

	input := make([]byte, 0)

	n := uint64(rand.Intn(10-2) + 2)
	k := uint64(rand.Intn(int(n)-2) + 2)

	_, err := DistributeBytes(input, n, k)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSecretSharingRecoverEmptyInput(t *testing.T) {
	restoreSlice := make([][]byte, 0)
	_, err := RecoverBytes(restoreSlice)
	if err == nil {
		log.Fatal("No error for empty input")
	}
}

func TestSecretSharingRecoverEmptyInput2(t *testing.T) {
	restoreSlice := make([][]byte, 100)
	_, err := RecoverBytes(restoreSlice)
	if err == nil {
		log.Fatal("No error for empty input")
	}
}

func TestSecretSharingInEqual(t *testing.T) {
	rand.Seed(time.Now().Unix())

	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n := uint64(rand.Intn(10-2) + 2)
	k := uint64(rand.Intn(int(n)-2) + 2)

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		log.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k-1)
	for i := uint64(0); i < k-1; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytes(restoreSlice)
	if err != nil {
		log.Fatal(err)
	}

	if bytes.Equal(input, nBytes) {
		log.Fatal("Input and output bytes are equal but they should not")
	}

}
