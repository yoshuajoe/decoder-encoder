package main

import (
	"fmt"
	"unsafe"
)

func PrintBinary(by byte) {
	for i := 7; i >= 0; i-- {
		b := by >> i
		if (b & 1) == 1 {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
}

func Encoder(number uint32) []byte {
	var temp []byte
	var i uintptr
	for i = 0; i < unsafe.Sizeof(number); i++ {
		pow := 8 * uint32(i)
		b := byte((number >> uint32(pow)) & 0xFF)
		temp = append(temp, b)
	}
	return temp
}

func Decoder(bytes []byte) uint32 {
	var result uint32
	var i uintptr
	result = uint32(bytes[0])
	for i = 1; i < unsafe.Sizeof(result); i++ {
		result += uint32(bytes[i]) * 256 * uint32(i)
	}
	return result
}

func main() {
	input := 2500
	bytes := Encoder(uint32(input))
	fmt.Printf("Encoding: %d\n", input)
	fmt.Print("[]bytes\t: ")
	for _, b := range bytes {
		fmt.Printf("%x ", b)
	}
	fmt.Printf("\n")

	fmt.Print("decimal\t: ")
	for _, b := range bytes {
		fmt.Printf("%d ", b)
	}
	fmt.Printf("\n")

	fmt.Print("binary\t: ")
	for _, b := range bytes {
		PrintBinary(b)
		fmt.Printf(" ")
	}
	fmt.Printf("\n")

	fmt.Printf("Decoding: %x\n", bytes)
	fmt.Printf("Decoded into: %d\n", Decoder(bytes))
}
