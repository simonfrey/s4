package main

import (
	"encoding/base64"
	"fmt"
	"github.com/simonfrey/s4"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
)

const version string = "0.5"
const AES_S4 string = "AES+S4"
const S4 string = "S4"

var formatRegex = regexp.MustCompile(`(?sm)====== BEGIN \[s4 v(\d+\.\d+) \|\| (AES\+S4|S4)\]======\n(.*?)\n====== END.*?======`)

var fmtString = func(useAES bool) string {
	opt := AES_S4
	if !useAES {
		opt = S4
	}
	return fmt.Sprintf("====== BEGIN [s4 v%s || %s]======\n%s\n====== END   [s4 v%s||%s]======\n",
		version, opt, "%s", version, opt)
}

func recoverShares(this js.Value, i []js.Value) interface{} {
	if len(i) != 1 {
		return js.Error{js.ValueOf(" input value required. Array of recover bytes (base64 encoded strings)")}
	}
	if i[0].Type() != js.TypeObject {
		return js.Error{js.ValueOf("First value must be the Input bytes (base64 encoded string)")}
	}

	inStrings := make([]string, i[0].Length())
	for k := 0; k < i[0].Length(); k++ {
		inStrings[k] = strings.TrimSpace(i[0].Index(k).String())
		if len(inStrings[k]) == 0 {
			return js.Error{js.ValueOf(fmt.Sprintf("Please provide all shares. Share '%d' is empty", k))}
		}
	}

	useAES := true
	format := ""

	shares := make([][]byte, len(inStrings))
	for k, v := range inStrings {
		if len(v) == 0 {
			continue
		}

		formatMatch := formatRegex.FindStringSubmatch(v)
		if len(formatMatch) != 4 {
			return js.Error{js.ValueOf("Invalid share input format for index " + strconv.Itoa(k))}
		}

		if format == "" {
			format = formatMatch[2]
			if format == S4 {
				useAES = false
			}
		} else if format != formatMatch[2] {
			return js.Error{js.ValueOf("Different formats in shares")}
		}

		tmpBytes, err := base64.StdEncoding.DecodeString(formatMatch[3])
		if err != nil {
			return js.Error{js.ValueOf("Could not decode base64 string: " + err.Error())}
		}
		shares[k] = tmpBytes
	}

	var clearText []byte
	var err error
	if useAES {
		clearText, err = s4.RecoverBytesAES(shares)
	} else {
		clearText, err = s4.RecoverBytes(shares)
	}
	if err != nil {
		return js.Error{js.ValueOf("Could not not recover bytes: " + err.Error())}
	}
	return js.ValueOf(base64.StdEncoding.EncodeToString(clearText))

}

func distributeShares(this js.Value, i []js.Value) interface{} {
	if len(i) < 3 {
		return js.Error{js.ValueOf("at least 3 input values required. Input bytes (base64 encoded string), n,k")}
	}
	if i[0].Type() != js.TypeString {
		return js.Error{js.ValueOf("First value must be the Input bytes (base64 encoded string)")}
	}

	if i[1].Type() != js.TypeNumber || i[2].Type() != js.TypeNumber {
		return js.Error{js.ValueOf("n,k must be of type number")}
	}

	if i[1].Int() < i[2].Int() {
		return js.Error{js.ValueOf("k must be smaller or equal to n")}
	}

	useAES := true
	if len(i) == 4 && i[3].Bool() == false {
		useAES = false
	}

	inBytes, err := base64.StdEncoding.DecodeString(i[0].String())
	if err != nil {
		return js.Error{js.ValueOf("Could not decode base64 string: " + err.Error())}
	}

	var byteShares [][]byte
	if useAES {
		byteShares, err = s4.DistributeBytesAES(inBytes, uint64(i[1].Int()), uint64(i[2].Int()))
	} else {
		byteShares, err = s4.DistributeBytes(inBytes, uint64(i[1].Int()), uint64(i[2].Int()))
	}
	if err != nil {
		return js.Error{js.ValueOf("Could not distribute bytes: " + err.Error())}
	}

	base64Shares := make([]interface{}, len(byteShares))
	for k, byteShare := range byteShares {
		base64Shares[k] = fmt.Sprintf(fmtString(useAES),
			base64.StdEncoding.EncodeToString(byteShare))
	}
	return js.ValueOf(base64Shares)
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
	// register functions
	registerCallbacks()
	<-c
}
