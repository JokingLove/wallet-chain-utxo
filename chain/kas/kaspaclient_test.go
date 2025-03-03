package kas

import (
	"testing"
)

func KaspaClient() *KasClient {
	client, err := NewKasClient("https://api.kaspa.org")
	if err != nil {
		panic(err)
	}
	return client
}

func Test_GetAccountBalance(t *testing.T) {
	address := "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
	balance, err := KaspaClient().GetAccountBalance(address)
	if err != nil {
		t.Error(err)
	}
	t.Log(balance)
}

func Test_GetUtxo(t *testing.T) {
	address := "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
	utxo, err := KaspaClient().GetUtxo(address)
	if err != nil {
		t.Error(err)
	}
	t.Log(utxo)
}

func Test_FeeEstimate(t *testing.T) {
	estimate, err := KaspaClient().FeeEstimate()
	if err != nil {
		t.Error(err)
	}
	t.Log(estimate)
}

func Test_GetBlockByHash(t *testing.T) {
	hash := "49fc8b62c9121b2ac3bce1e15f107f31ca67b66b2f24733a4722b3e23f7db31f"
	block, err := KaspaClient().GetBlockByHash(hash)
	if err != nil {
		t.Error(err)
	}
	t.Log(block)
}

func Test_GetBlockByHeight(t *testing.T) {
	height, err := KaspaClient().GetBlockByHeight(103082222)
	if err != nil {
		t.Error(err)
	}
	t.Log(height)
}

func Test_GetBlockByHeight0(t *testing.T) {
	height, err := KaspaClient().GetBlockByHeight(0)
	if err != nil {
		t.Error(err)
	}
	t.Log(height)
}

func Test_GetTxByHash(t *testing.T) {
	hash := "5387f2e272236771845f4bf0a5f0b0f41daaed1846f17ee266d3a50f616e8d59"
	transaction, err := KaspaClient().GetTxByHash(hash)
	if err != nil {
		t.Error(err)
	}
	t.Log(transaction)
}

func Test_GetTxByAddress(t *testing.T) {
	address := "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
	transaction, err := KaspaClient().GetTxByAddress(address)
	if err != nil {
		t.Error(err)
	}
	t.Log(transaction)
}
