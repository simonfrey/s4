package s4

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randNK() (n uint64, k uint64) {
	max := 255
	min := 2
	nI := rand.Intn(max-min) + min
	kI := 2
	if nI-min > 0 {
		kI = rand.Intn(nI-min) + min
	}
	return uint64(nI), uint64(kI)
}

// Normal
func TestSecretSharingEqual(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytes(restoreSlice)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(input, nBytes) {
		t.Fatal("Input and output bytes differ")
	}

}

func TestSecretSharingSameShares(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		restoreSlice[i] = shares[0]
	}

	_, err = RecoverBytes(restoreSlice)
	if err == nil {
		t.Fatal("Duplicated parts where accepted")
	}
}

func TestSecretSharingDistributeEmptyInput(t *testing.T) {
	input := make([]byte, 0)

	n, k := randNK()

	_, err := DistributeBytes(input, n, k)
	if err == nil {
		t.Fatal("Tried to distribute empty input")
	}
}

func TestSecretSharingRecoverEmptyInput(t *testing.T) {
	restoreSlice := make([][]byte, 0)
	_, err := RecoverBytes(restoreSlice)
	if err == nil {
		t.Fatal("No error for empty input")
	}
}

func TestSecretSharingRecoverEmptyInput2(t *testing.T) {
	restoreSlice := make([][]byte, 100)
	_, err := RecoverBytes(restoreSlice)
	if err == nil {
		t.Fatal("No error for empty input")
	}
}

func TestSecretSharingInEqual(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytes(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k-1)
	for i := uint64(0); i < k-1; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytes(restoreSlice)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(input, nBytes) {
		t.Fatal("Input and output bytes are equal but they should not")
	}

}

// AES
func TestSecretSharingEqualAES(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytesAES(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytesAES(restoreSlice)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(input, nBytes) {
		t.Fatal("Input and output bytes differ")
	}

}

func TestSecretSharingSameSharesAES(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytesAES(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	restoreSlice := make([][]byte, k)
	for i := uint64(0); i < k; i++ {
		restoreSlice[i] = shares[0]
	}

	_, err = RecoverBytesAES(restoreSlice)
	if err == nil {
		t.Fatal("Duplicated parts where accepted")
	}
}

func TestSecretSharingDistributeEmptyInputAES(t *testing.T) {
	input := make([]byte, 0)

	n, k := randNK()

	_, err := DistributeBytesAES(input, n, k)
	if err == nil {
		t.Fatal("Tried to distribute empty input")
	}
}

func TestSecretSharingRecoverEmptyInputAES(t *testing.T) {
	restoreSlice := make([][]byte, 0)
	_, err := RecoverBytesAES(restoreSlice)
	if err == nil {
		t.Fatal("No error for empty input")
	}
}

func TestSecretSharingRecoverEmptyInput2AES(t *testing.T) {
	restoreSlice := make([][]byte, 100)
	_, err := RecoverBytesAES(restoreSlice)
	if err == nil {
		t.Fatal("No error for empty input")
	}
}

func TestSecretSharingInEqualAES(t *testing.T) {
	input := make([]byte, rand.Intn(30000-1000)+1000)
	rand.Read(input)

	n, k := randNK()

	shares, err := DistributeBytesAES(input, n, k)
	if err != nil {
		t.Fatal(err)
	}

	shift := uint64(rand.Intn(100-20) + 20)
	restoreSlice := make([][]byte, k-1)
	for i := uint64(0); i < k-1; i++ {
		ik := (i + shift) % n
		restoreSlice[i] = shares[ik]
	}

	nBytes, err := RecoverBytesAES(restoreSlice)
	if err == nil {
		t.Fatal("Recovery worked, but it should not")

	}

	if bytes.Equal(input, nBytes) {
		t.Fatal("Input and output bytes are equal but they should not")
	}

}
