package blockchain

import "testing"

func TestWallet(t *testing.T) {
	t.Log(string(NewWallet().GetAddress()))
}
