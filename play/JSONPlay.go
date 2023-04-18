package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"rvctomqtt/can"
	"rvctomqtt/constants"
	"rvctomqtt/convert"
	"rvctomqtt/utils"
	"time"
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

type hexUint32 uint32
type hexUint16 uint16
type hexUint8 uint8

func (d hexUint32) MarshalYAML() (interface{}, error) {
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

func pointer(x byte) *byte {
	return &x
}

// We can call a parser function based on data type and size

//func getByte(data []byte, idx int)

type rvcmessageField struct {
	Name   string
	DgnHI  hexUint16
	DgnLOW hexUint8
	Fields map[string]rvcDataField
}
type rvcDataField struct {
	//Name         string
	BytePosition  byte
	BitPosition   *byte `yaml:",omitempty"` // omits 0 if we don't use a pointer - and we need to differentiate between 0 and not there
	DataType      dataFieldType
	DataValueType dataValueType `yaml:",omitempty"`
	// ValueType     string          // Degrees C, Percent, Volts, etc. Combination of  DataType and ValueType tell us how to parse
	Values map[byte]string `yaml:",omitempty"`
	//BitsSize     byte
}

type dataFieldType string
type dataValueType string

type dataFieldTypeSS struct {
}

const (
	Bitfield2      dataFieldType = "bit2"
	CharacterField               = "char8"
	Uint8Field                   = "uint8"
	Uint16Field                  = "uint16"
	Uint32Field                  = "uint32"

	Percent    dataValueType = "%"
	Instance                 = "instance"
	DegreesC8                = "Degrees C8"
	DegreesC16               = "Degrees C16"
	Volts8                   = "V8"
	Volts16                  = "V16"
	Amps8                    = "A8"
	Amps16                   = "A16"
	Amps32                   = "A32"
	Hz                       = "Hz"
	Watts                    = "Watts"
	AmpHours                 = "AmpHours"
)

func main2() {

	//   DC current:
	//        byteposition: 4
	//        datatype: uint32
	//        datavaluetype: A

	// DGN (hi/Low)
	// Name
	// Fields
	//    Name, Value, valueMap (Known Values )

	var mf2 = map[hexUint32]rvcmessageField{
		0x1FFFD: {
			Name:   "DC_SOURCE_STATUS_1",
			DgnHI:  0x1FF,
			DgnLOW: 0xFD,
			Fields: map[string]rvcDataField{
				"instance": {
					BytePosition:  0,
					DataType:      Uint8Field,
					DataValueType: Instance,
					Values: map[byte]string{
						0: "invalid",
						1: "main house",
						2: "chassis",
						3: "secondary house",
					},
				},
				"device priority": {
					BytePosition: 1,
					DataType:     Uint8Field,
					Values: map[byte]string{
						120: "battery soc device",
						100: "inverter Charger",
						80:  "charger",
						60:  "inverter",
						40:  "voltmeter ammeter",
						20:  "voltmeter",
					},
				},
				"DC voltage": {
					BytePosition:  2,
					DataType:      Uint16Field,
					DataValueType: Volts16,
				},
				"DC current": {
					BytePosition:  4,
					DataType:      Uint32Field,
					DataValueType: Amps32,
				},
			},
		},
	}

	//x := [5]int{10, 20, 30, 40, 50}   // Intialized with values

	//var tstData = [8]byte{
	//	0x01, 0x64, 0x0c, 0x01, 0xa0, 0x26, 0x35, 0x77,
	//}

	//var bytePos = mf2[0x1FFFD].Fields["instance"].BytePosition
	//var dt = mf2[0x1FFFD].Fields["instance"].DataType
	//var bp = mf2[0x1FFFD].Fields["instance"].BitPosition

	//switch dt {
	//case Amps32:
	//	var val = utils.Getuint32(&tstData, bytePos)
	//	var fval = convert.ToCurrent(val)
	//	break
	//case Amps16:
	//case Volts16:
	//
	//	break
	//default:
	//
	//}

	//var mf []rvcmessageField = []rvcmessageField{
	//	{
	//		Name:   "DC_SOURCE_STATUS_1",
	//		DgnHI:  0x1FF,
	//		DgnLOW: 0xFD,
	//		Fields: map[string]rvcDataField{
	//			"instance": {
	//				BytePosition:  0,
	//				DataType:      Uint8Field,
	//				DataValueType: Instance,
	//				Values: map[byte]string{
	//					0: "invalid",
	//					1: "main house",
	//					2: "chassis",
	//					3: "secondary house",
	//				},
	//			},
	//			"device priority": {
	//				BytePosition: 1,
	//				DataType:     Uint8Field,
	//				Values: map[byte]string{
	//					120: "battery soc device",
	//					100: "inverter Charger",
	//					80:  "charger",
	//					60:  "inverter",
	//					40:  "voltmeter ammeter",
	//					20:  "voltmeter",
	//				},
	//			},
	//			"DC voltage": {
	//				BytePosition:  2,
	//				DataType:      Uint16Field,
	//				DataValueType: Volts,
	//				// NEED DATA TYPE DEGc
	//			},
	//			"DC current": {
	//				BytePosition:  4,
	//				DataType:      Uint32Field,
	//				DataValueType: Amps,
	//			},
	//		},
	//	},
	//	{
	//		Name:   "DGN2",
	//		DgnHI:  0x100,
	//		DgnLOW: 0x22,
	//		Fields: map[string]rvcDataField{
	//			"field1": {
	//				//Name:         "field1x",
	//				BytePosition: 0,
	//				BitPosition:  pointer(1),
	//				DataType:     Bitfield,
	//			},
	//			"field2": {
	//				//Name:         "field2x",
	//				BytePosition: 0,
	//				DataType:     Uint16Field,
	//			},
	//			"field3": {
	//				//Name:         "field3x",
	//				BytePosition: 1,
	//				DataType:     Uint32Field,
	//			},
	//		},
	//	},
	//}

	var vv, err = json.Marshal(&mf2)
	if err != nil {
		log.Fatal(err)
		return
	}
	var bb bytes.Buffer
	json.Indent(&bb, vv, "", "\t")
	fmt.Println("------------------ JSON ----------------------------")
	bb.WriteTo(os.Stdout)
	fmt.Println("------------------ END JSON -------------------------")

	d, err := yaml.Marshal(&mf2)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	fmt.Println("------------------ YAML ----------------------------")
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	fmt.Println("------------------ END YAML -------------------------")

	//d2, err2 := yaml.Marshal(&mf2)
	//if err2 != nil {
	//	log.Fatalf("error: %v", err2)
	//	return
	//}
	//fmt.Println("------------------ YAML2 ----------------------------")
	//fmt.Printf("--- t dump:\n%s\n\n", string(d2))
	//fmt.Println("------------------ END YAML2 -------------------------")

	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	//var decoded = [2]rvcmessageField{}
	var decoded = map[hexUint32]rvcmessageField{}

	var error = yaml.Unmarshal(d, &decoded)

	if error != nil {
		log.Fatalf("error: %v", error)
		return
	}
	fmt.Printf("--- t decoded:\n%x\n\n", decoded)
	//fmt.Printf("--- t dump:\n%s\n\n", decoded)

	r, e := yaml.Marshal(&decoded)
	if e != nil {
		log.Fatalf("error: %v", e)
		return
	}
	fmt.Println("------------------ YAML - Decoded and encoded ----------------------------")
	fmt.Printf("--- t dump:\n%s\n\n", string(r))
	fmt.Println("------------------ END YAML -------------------------")

	var ts = time.Now()
	var dataBytes uint8 = 8
	var cf = can.Frame{
		Timestamp: ts,
		ID:        0,
		Length:    0,
		Flags:     0,
		Res0:      0,
		Res1:      0,
		Data:      [constants.MaxFrameDataLength]uint8{},
		// for a test we want the 1st for raw bytes to be BACKWARDS - MS Byte first
		// if all is well the ID will end up with the bits flipped.
		// It's possible that this test may fail due to a platform change
		// 0 = 0x77359400
		MessageBytes: [constants.MAX_MESSAGE]byte{1, 0, 0, 0xE0, dataBytes, 0, 0, 0, 0x00, 0x01, 0x02, 0x03, 0x00, 0x94, 0x35, 0x77},
	}

	cf.BuildCanFrameX()
	//	DC current:
	//byteposition: 4
	//datatype: uint32
	//datavaluetype: A

	//func Getuint32(buf *[constants.MaxFrameDataLength]byte, byteOffset byte) uint32 {

	var converted = convert.ToCurrent(utils.Getuint32(&cf.Data, 4))
	//var p = utils.UintParser{ByteOffset: 4}.ParseInt32
	//var conv = convert.ToCurrent
	//var converted = conv(p(&cf.Data))
	//
	fmt.Printf("Converted  %f", converted)

}

type bit2 uint8

type Number interface {
	uint32 | uint16 | uint8 | float64 | bit2
}

type field[T Number] struct {
	name  string
	value T
}

func (f *field[_]) GetName() string {
	return f.name
}
func (f *field[Number]) getValue() Number {
	return f.value
}

type thingA struct {
	field_A uint8
	field_B uint16
	field_C uint32
	field_D float64
	field_E bit2
}

func main() {
	//var f1 = field[uint8]{name: "f1", value: 2}
	//var f2 = field[uint16]{name: "f2", value: 2}
	//var f3 = field[uint32]{name: "f3", value: 2}
	//var f4 = field[float64]{name: "f4", value: 2}
	//var f5 = field[bit2]{name: "f5", value: 2}
	//var ff = []any{f1, f2, f3, f4, f5}

	var tt = thingA{}

	//for _, v := range ff {
	//	var vv, ok = v.(field[uint8])

	s := reflect.ValueOf(&tt).Elem()
	typeOfT := s.Type()
	var k = s.Kind()
	fmt.Printf("kind = %s\n", k)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i,
		//	typeOfT.Field(i).Name, f.Type(), f.Interface())

		fmt.Printf("%d: %s %s \n", i,
			typeOfT.Field(i).Name, f.Type())
	}
	//}
	fmt.Printf("Hey")
}
