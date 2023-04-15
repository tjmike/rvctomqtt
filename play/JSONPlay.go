package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

//
//"messageSpec": {
//"DGNName": {
//"dgnHi": "0x111",
//"dgnLow": "0x11",
//"fields": {
//"Field1": {
//"byteOffset": 0,
//"bitOffset": 0,
//"bits": 8,
//"status": "retired",
//"engine": "Gecko",
//"engine_version": "1.7"
//}
//}
//}
//}
//}

type hexUint16 uint16
type hexUint8 uint8

func (d hexUint16) MarshalYAML() (interface{}, error) {
	var hx = fmt.Sprintf("0X%x", d)
	node := yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FlowStyle,
		Value: hx,
	}
	// This comment WILL be included in the output
	//node.HeadComment = "HEX DATA MarshalYAML Data"
	return node, nil
}
func (d hexUint8) MarshalYAML() (interface{}, error) {
	var hx = fmt.Sprintf("0X%x", d)
	node := yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FlowStyle,
		Value: hx,
	}
	// This comment WILL be included in the output
	//node.HeadComment = "HEX DATA MarshalYAML Data"
	return node, nil
}

type rvcmessageField struct {
	Name   string
	DgnHI  hexUint16
	DgnLOW hexUint8
	Fields map[string]rvcDataField
}
type rvcDataField struct {
	//Name         string
	BytePosition byte
	BitPosition  byte
	DataType     string
	Values       map[byte]string `yaml:",omitempty"`
	//BitsSize     byte
}

const (
	Bitfield       string = "Bit2"
	CharacterField        = "char8"
	Uint8Field            = "uint8"
	Uint16Field           = "uint16"
	Uint32Field           = "uint32"
)

func main() {
	var mf []rvcmessageField = []rvcmessageField{
		{
			Name:   "DGN1",
			DgnHI:  0x1ff,
			DgnLOW: 0x11,
			Fields: map[string]rvcDataField{
				"instance": {
					//Name:         "field1",
					BytePosition: 0,
					//BitPosition:  0,
					DataType: Uint32Field,
					Values: map[byte]string{
						1: "instance 1",
						2: "instance 2",
						3: "instance 2",
					},
				},
				"field2": {
					//Name:         "field2",
					BytePosition: 4,
					BitPosition:  1,
					DataType:     Bitfield,
				},
				"field3": {
					//Name:         "field3",
					BytePosition: 4,
					BitPosition:  2,
					DataType:     Bitfield,
				},
			},
		},
		{
			Name:   "DGN2",
			DgnHI:  0x100,
			DgnLOW: 0x22,
			Fields: map[string]rvcDataField{
				"field1": {
					//Name:         "field1x",
					BytePosition: 0,
					DataType:     Bitfield,
				},
				"field2": {
					//Name:         "field2x",
					BytePosition: 0,
					DataType:     Uint16Field,
				},
				"field3": {
					//Name:         "field3x",
					BytePosition: 1,
					DataType:     Uint32Field,
				},
			},
		},
	}

	var vv, err = json.Marshal(&mf)
	if err != nil {
		log.Fatal(err)
		return
	}
	var bb bytes.Buffer
	json.Indent(&bb, vv, "", "\t")
	fmt.Println("------------------ JSON ----------------------------")
	bb.WriteTo(os.Stdout)
	fmt.Println("------------------ END JSON -------------------------")

	d, err := yaml.Marshal(&mf)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	fmt.Println("------------------ YAML ----------------------------")
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	fmt.Println("------------------ END YAML -------------------------")

	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	var decoded = [2]rvcmessageField{}

	var error = yaml.Unmarshal(d, &decoded)

	if error != nil {
		log.Fatalf("error: %v", error)
		return
	}
	fmt.Printf("--- t decoded:\n%x\n\n", decoded)
	fmt.Printf("--- t dump:\n%s\n\n", decoded)

}
