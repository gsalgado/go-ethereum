package abi

import (
	"fmt"
	"reflect"
)

// This is meant to be used only on tests, but it's exported because otherwise
// tests would have to duplicate a lot of code from this package (e.g.
// Type.pack(), Type.requiresLengthPrefix(), packNum()).
func PackMockResponse(m Method, outputs ...interface{}) ([]byte, error) {
	// Make sure arguments match up and pack them
	if len(outputs) != len(m.Outputs) {
		return nil, fmt.Errorf("output count mismatch: %d for %d", len(outputs), len(m.Outputs))
	}
	// variable output is the output appended at the end of packed
	// output. This is used for strings and bytes types output.
	var variableOutput []byte

	var ret []byte
	for i, o := range outputs {
		output := m.Outputs[i]
		packed, err := output.Type.pack(reflect.ValueOf(o))
		if err != nil {
			return nil, fmt.Errorf("`%s` %v", m.Name, err)
		}

		// check for a slice type (string, bytes, slice)
		if output.Type.requiresLengthPrefix() {
			// calculate the offset
			offset := len(m.Outputs)*32 + len(variableOutput)
			// set the offset
			ret = append(ret, packNum(reflect.ValueOf(offset))...)
			// Append the packed output to the variable output. The variable
			// output will be appended at the end of the output.
			variableOutput = append(variableOutput, packed...)
		} else {
			// append the packed value to the output
			ret = append(ret, packed...)
		}
	}
	// append the variable output at the end of the packed output
	ret = append(ret, variableOutput...)

	return ret, nil
}
