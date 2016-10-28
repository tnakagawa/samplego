// hex project main.go
package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	// Decode
	bs, _ := hex.DecodeString("0A1A2A3A4A")
	fmt.Printf("%02x\n", bs)
	//  Encode
	str := hex.EncodeToString(bs)
	fmt.Printf("%s\n", str)

	bs = make([]byte, 32)
	fmt.Printf("%02x\n", bs)
	str = hex.EncodeToString(bs)
	fmt.Printf("%s\n", str)

	i := uint32(0)
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	str = hex.EncodeToString(buf)
	fmt.Printf("%s\n", str)
}
