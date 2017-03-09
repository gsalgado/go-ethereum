package mock

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MockMethod creates a method with the given name on the given Proxy, using
// the same arguments as the method with the same name from the given contract
// ABI. The mock method will always return the values defined by outputs.
func MockMethod(p *Proxy, name string, outputs []interface{}, contract string, opts *bind.TransactOpts) error {
	contractABI, err := abi.JSON(strings.NewReader(contract))
	if err != nil {
		return err
	}
	abiMethod, exists := contractABI.Methods[name]
	if !exists {
		return fmt.Errorf("method %s not found", name)
	}
	params := make([]interface{}, len(abiMethod.Inputs))
	for i, p := range abiMethod.Inputs {
		params[i] = p
	}
	input, err := contractABI.Pack(name, params...)
	if err != nil {
		return err
	}
	mockResponse, err := abi.PackMockResponse(abiMethod, outputs...)
	if err != nil {
		return err
	}
	_, err = p.SetMockWithArgs(opts, input, 2, common.Address{}, mockResponse)
	return err
}
