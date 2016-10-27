// hex project main.go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	bs, _ := hex.DecodeString("0A1A2A3A4A")
	fmt.Printf("%02x\n", bs)
	str := hex.EncodeToString(bs)
	fmt.Printf("%s\n", str)

	bs = make([]byte, 32)
	fmt.Printf("%02x\n", bs)
	str = hex.EncodeToString(bs)
	fmt.Printf("%s\n", str)
}
