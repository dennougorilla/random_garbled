package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bytes := []byte{0xbd, 0xb2, 0x3d, 0xbc, 0x20, 0xe2, 0x8c, 0x98}
	s := string(bytes)
	fmt.Printf("% x\n", s)
	fmt.Println(s)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	fmt.Printf("% x\n", sample)

	src := int64(123456789123456789)
	result := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(result, src)
	fmt.Println(result)
}

//print(int(random.randrange(0,65535)).to_bytes(2,byteorder='little').decode('shift-jis','replace'), end='')
