package main

import (
	"encoding/base64"
	"fmt"
	"github.com/simonfrey/s4/internal/shares_logic"
	"strings"
	"syscall/js"
)

func jsError(f string, a ...any) string {
	return fmt.Errorf(f, a...).Error()
}

// WASM Stubs
func recoverShares(this js.Value, i []js.Value) interface{} {
	if len(i) != 1 {
		return jsError("input value required. Array of recover bytes (base64 encoded strings)")
	}
	if i[0].Type() != js.TypeObject {
		return jsError("First value must be the Input bytes (base64 encoded string)")
	}

	inStrings := make([]string, i[0].Length())
	for k := 0; k < i[0].Length(); k++ {
		inStrings[k] = strings.TrimSpace(i[0].Index(k).String())
	}

	result, err := shares_logic.RecoverShares(inStrings)
	if err != nil {
		return jsError("could not recover shares: %w", err)
	}
	base64Result := base64.StdEncoding.EncodeToString(result)

	return base64Result
}

func distributeShares(this js.Value, i []js.Value) interface{} {
	if len(i) < 3 {
		return jsError("at least 3 input values required. Input bytes (base64 encoded string), n,k")
	}
	if i[0].Type() != js.TypeString {
		return jsError("First value must be the Input bytes (base64 encoded string)")
	}
	if i[1].Type() != js.TypeNumber || i[2].Type() != js.TypeNumber {
		return jsError("n,k must be of type number")
	}
	if i[1].Int() < i[2].Int() {
		return jsError("k must be smaller or equal to n")
	}

	useAES := true
	if len(i) >= 4 && i[3].Bool() == false {
		useAES = false
	}
	useBase24 := true
	if len(i) >= 5 && i[4].Bool() == false {
		useBase24 = false
	}

	inputBytes, err := base64.StdEncoding.DecodeString(i[0].String())
	if err != nil {
		return jsError("could not decode base64 string of input: %w", err)
	}

	result, err := shares_logic.DistributeShares(inputBytes, i[1].Int(), i[2].Int(), useAES, useBase24)
	if err != nil {
		return jsError("could not distribute shares: %w", err)
	}

	jsShares := make([]interface{}, len(result))
	for k, v := range result {
		jsShares[k] = v
	}
	return jsShares
}

func registerCallbacks() {
	js.Global().Set("Distribute_fours", js.FuncOf(distributeShares))
	js.Global().Set("Recover_fours", js.FuncOf(recoverShares))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
