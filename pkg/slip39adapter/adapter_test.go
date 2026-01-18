package slip39adapter

import "testing"

func TestRoundTrip(t *testing.T) {
	share := "[s4 v0.6 aes+s4 example-share-data]"

	mn, err := ShareToSlip39Mnemonic(share)
	if err != nil {
		t.Fatalf("unexpected error converting to mnemonic: %v", err)
	}

	back, err := Slip39MnemonicToShare(mn)
	if err != nil {
		t.Fatalf("unexpected error converting back: %v", err)
	}

	if back != share {
		t.Fatalf("roundtrip mismatch: want %q got %q", share, back)

	}
}
