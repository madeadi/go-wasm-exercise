package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Prettify Json!")
	js.Global().Set("prettyJson", jsonWrapper())
	<-make(chan bool)
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments"
		}

		inputJson := args[0].String()
		fmt.Printf("Input: %s\n", inputJson)
		pretty, err := prettyJson(inputJson)
		if err != nil {
			fmt.Printf("Unable to convert json: %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
