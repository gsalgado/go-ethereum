// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mock

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"method\",\"type\":\"bytes4\"},{\"name\":\"operationType\",\"type\":\"uint8\"},{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setMock\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"methodWithData\",\"type\":\"bytes\"},{\"name\":\"operationType\",\"type\":\"uint8\"},{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setMockWithArgs\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"traceFunctionCalls\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Trace\",\"type\":\"event\"}]"

// ProxyBin is the compiled bytecode used for deploying new contracts.
const ProxyBin = `60606040526000600060006101000a81548160ff021916908315150217905550341561002757fe5b604051602080610a60833981016040528080519060200190919050505b80600060006101000a81548160ff0219169083151502179055505b505b6109f0806100706000396000f3006060604052361561004a576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806361d860f4146104ff57806362df51f4146105ac575b341561005257fe5b6104fd5b61005e6108b6565b60006000600061006c6108f7565b600060009054906101000a900460ff1615610110577ffb5a7aca965dae444c2cb92e2fac23a21279542c04537a28534e1a4ed598ca393360003634604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825285858281815260200192508082843782019150509550505050505060405180910390a15b6002600036604051808383808284378201915050925050509081526020016040518091039020606060405190810160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b50505050508152505094508460000151935060008460ff16141561041d57600160006000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103f95780601f106103ce576101008083540402835291602001916103f9565b820191906000526020600020905b8154815290600101906020018083116103dc57829003601f168201915b50505050508152505094508460000151935060008460ff16141561041c57610000565b5b60018460ff1614156104d457846020015192508273ffffffffffffffffffffffffffffffffffffffff16856040015160405180828051906020019080838360008314610488575b80518252602083111561048857602082019150602081019050602083039250610464565b505050905090810190601f1680156104b45780820380516001836020036101000a031916815260200191505b509150506000604051808303816000866161da5a03f191505091506104f5565b60028460ff1614156104ef5784604001519050805160208201f35b610000565b5b5b5050505050565b005b341561050757fe5b6105aa60048080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909190803560ff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610674565b005b34156105b457fe5b610672600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803560ff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610789565b005b6060604051908101604052808460ff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281525060016000867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600101908051906020019061077e92919061090b565b509050505b50505050565b6060604051908101604052808460ff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152506002856040518082805190602001908083835b602083106107f457805182526020820191506020810190506020830392506107d1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010190805190602001906108ab92919061090b565b509050505b50505050565b606060405190810160405280600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016108f161098b565b81525090565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061094c57805160ff191683800117855561097a565b8280016001018555821561097a579182015b8281111561097957825182559160200191906001019061095e565b5b509050610987919061099f565b5090565b602060405190810160405280600081525090565b6109c191905b808211156109bd5760008160009055506001016109a5565b5090565b905600a165627a7a72305820e644b8e1556e1b38f377c011e9a7dd4a53761ce58eed85cea3421cba63da02960029`

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, traceFunctionCalls bool) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend, traceFunctionCalls)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// SetMock is a paid mutator transaction binding the contract method 0x61d860f4.
//
// Solidity: function setMock(method bytes4, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxyTransactor) SetMock(opts *bind.TransactOpts, method [4]byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setMock", method, operationType, target, data)
}

// SetMock is a paid mutator transaction binding the contract method 0x61d860f4.
//
// Solidity: function setMock(method bytes4, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxySession) SetMock(method [4]byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetMock(&_Proxy.TransactOpts, method, operationType, target, data)
}

// SetMock is a paid mutator transaction binding the contract method 0x61d860f4.
//
// Solidity: function setMock(method bytes4, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxyTransactorSession) SetMock(method [4]byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetMock(&_Proxy.TransactOpts, method, operationType, target, data)
}

// SetMockWithArgs is a paid mutator transaction binding the contract method 0x62df51f4.
//
// Solidity: function setMockWithArgs(methodWithData bytes, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxyTransactor) SetMockWithArgs(opts *bind.TransactOpts, methodWithData []byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setMockWithArgs", methodWithData, operationType, target, data)
}

// SetMockWithArgs is a paid mutator transaction binding the contract method 0x62df51f4.
//
// Solidity: function setMockWithArgs(methodWithData bytes, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxySession) SetMockWithArgs(methodWithData []byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetMockWithArgs(&_Proxy.TransactOpts, methodWithData, operationType, target, data)
}

// SetMockWithArgs is a paid mutator transaction binding the contract method 0x62df51f4.
//
// Solidity: function setMockWithArgs(methodWithData bytes, operationType uint8, target address, data bytes) returns()
func (_Proxy *ProxyTransactorSession) SetMockWithArgs(methodWithData []byte, operationType uint8, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetMockWithArgs(&_Proxy.TransactOpts, methodWithData, operationType, target, data)
}
