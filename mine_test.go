package main

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
)

func TestEncoder(t *testing.T) {
	sampleVar := uint32(2500)
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, uint32(sampleVar))

	expected := buff.Bytes()
	result := Encoder(sampleVar)

	if !reflect.DeepEqual(expected, result) {
		t.Fatal("not tallied")
	}
}

func TestDecoder(t *testing.T) {
	sampleVar := uint32(2500)
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, uint32(sampleVar))

	expected := uint32(2500)
	result := Decoder(buff.Bytes())

	if expected != result {
		t.Fatalf("not tallied expected: %d | result: %d", expected, result)
	}
}
