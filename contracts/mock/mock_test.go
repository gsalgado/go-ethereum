package mock

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/release"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
	opts   = bind.NewKeyedTransactor(key)
)

func TestMockMethod(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAccount{Address: addr, Balance: big.NewInt(1000000000)})
	proxyAddr, _, proxy, err := DeployProxy(opts, backend, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	releaseContract, err := release.NewReleaseOracle(proxyAddr, backend)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	signerKey, _ := crypto.GenerateKey()
	signerAddr := crypto.PubkeyToAddress(signerKey.PublicKey)
	err = MockMethod(proxy, "signers", []interface{}{[]common.Address{signerAddr}}, release.ReleaseOracleABI, opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	backend.Commit()

	signers, err := releaseContract.Signers(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(signers) != 1 {
		t.Fatalf("expected 1 signer, got %d", len(signers))
	}
	if signers[0].Hex() != signerAddr.Hex() {
		t.Fatalf("expected signer to be %s, got %s", signerAddr.Hex(), signers[0].Hex())
	}
}

func TestMockMethodNotFound(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAccount{Address: addr, Balance: big.NewInt(1000000000)})
	_, _, proxy, err := DeployProxy(opts, backend, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = MockMethod(proxy, "nonexistent", []interface{}{}, release.ReleaseOracleABI, opts)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
