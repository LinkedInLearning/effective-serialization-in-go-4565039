package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	printEncoding(int64(12345678))
	printEncoding(2.718281)
	printEncoding("I ♡ Go")
}

func printEncoding(val any) {
	v := val
	if s, ok := v.(string); ok {
		v = []byte(s)
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
		fmt.Printf("%#v: error - %s\n", val, err)
		return
	}

	fmt.Printf("%#v: %x\n", val, buf.Bytes())
}
