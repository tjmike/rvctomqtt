package rvc

import (
	"fmt"
	"testing"
)

func TestAddressClaim2(t *testing.T) {
	var ac = InformationRequest{}
	var f = RvcFrame{}
	var raw = f.GetMessage()

	var desiredSA uint8 = 0x14
	ac.SetAsAddressClaim(desiredSA)
	// DGN 0xEA00 -> 0xEA14 (request address 0x14)

	// pri 3 res 1 page 1
	//    PRI(3)
	//        RES(1)
	//         DGN HIGH    DGNLOW    SA
	// 0000 0000 0000 0000 0000 0000 0000 0000
	// 1001 1000 E    A    1    4    F    E
	// 9    8
	//
	// 98 ea 14 fe

	// 9f db fe 99 08 00 00 00 3b ff c8 01 02 00 ff ff
	raw[0] = 0xfe
	raw[1] = desiredSA
	raw[2] = 0xea
	raw[3] = 0x98
	raw[4] = 0x08
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0
	raw[8] = 0x00
	raw[9] = 238
	raw[10] = 0x00
	raw[11] = 0xFF
	raw[12] = 0xFF
	raw[13] = 0xFF
	raw[14] = 0xFF
	raw[15] = 0xff

	f.BuildCanFrameX()
	ac.Init(&f)

	if ac.GetDestinationAddress() != 0x14 {
		t.Errorf("Wrong requested source address  expected: %x got %x\n", desiredSA, ac.SourceAddress)
	}

	if !ac.isAddressClaim {
		t.Errorf("Expected addressClaimed to be set\n")

	}
	//fmt.Printf("raw = %s\n", raw)
	// why is this print needed to get the test to run???
	fmt.Printf("Addr Clm: %s\n", ac.String())
}

func TestAddressClaimCreateFrame2(t *testing.T) {

	var acOrig = InformationRequest{}
	acOrig.DGN = DGN_INFORMATION_REQUEST
	acOrig.Name = DGNName(acOrig.DGN)
	acOrig.SetPriority(0x06)
	acOrig.SourceAddress = 0xFE

	// set AFTER setting up the DGN so it can be update correctly
	acOrig.SetAsAddressClaim(0x14)

	var createdFrame = acOrig.CreateFrame()

	var dataBytesFrame = RvcFrame{}
	var bytes = dataBytesFrame.GetMessage()
	for i := 0; i < len(bytes); i++ {
		dataBytesFrame.GetMessage()[i] = createdFrame.GetMessage()[i]
	}

	dataBytesFrame.BuildCanFrameX()
	var rebuiltAddressClaim = InformationRequest{}
	rebuiltAddressClaim.Init(&dataBytesFrame)

	fmt.Printf("OUT = %x\n", createdFrame.MessageBytes)
	fmt.Printf("OUT = %x\n", dataBytesFrame.MessageBytes)

	if !acOrig.Equals(&rebuiltAddressClaim) {
		t.Errorf("Initial and reconstructed address claim commands dont match %s != %s\n", acOrig.String(), rebuiltAddressClaim.String())
	}

}

func TestAddressClaimCreateFrame3(t *testing.T) {

	var destAddr uint8 = 0x01
	var desiredDGN = DGN_PRODUCT_IDENTIFICATION_MESSAGE

	var acOrig = InformationRequest{}
	acOrig.SetDGN(DGN_INFORMATION_REQUEST)
	acOrig.Name = DGNName(acOrig.DGN)
	acOrig.Priority = 0x06
	acOrig.SourceAddress = 0xFE

	acOrig.SetDestinationAddress(destAddr)
	acOrig.SetDesiredDGN(desiredDGN)

	var createdFrame = acOrig.CreateFrame()

	var dataBytesFrame = RvcFrame{}
	var bytes = dataBytesFrame.GetMessage()
	for i := 0; i < len(bytes); i++ {
		dataBytesFrame.GetMessage()[i] = createdFrame.GetMessage()[i]
	}

	dataBytesFrame.BuildCanFrameX()
	var rebuilt = InformationRequest{}
	rebuilt.Init(&dataBytesFrame)

	if rebuilt.GetDesiredDGN() != desiredDGN {
		t.Errorf("Desired DGN was wrong expected %x got %x\n",
			desiredDGN, rebuilt.GetDesiredDGN())

	}
	if !acOrig.Equals(&rebuilt) {
		t.Errorf("Initial and reconstructed address claim commands dont match %s != %s\n", acOrig.String(), rebuilt.String())
	}

}
