package rvc

import (
	"encoding/binary"
	"fmt"
	"rvctomqtt/constants"
	"testing"
)

func TestAddressClaimed(t *testing.T) {
	var claimed = AddressClaimed{}
	var f = RvcFrame{}
	var raw = f.GetMessage()

	// // bit 0-28: CAN identifier (11/29 bit)
	//	// bit 29: error message flag (ERR)
	//	// bit 30: remote transmision request (RTR)
	//	// bit 31: extended frame format (EFF) (MSB)
	//
	// 29 bit can frame
	// 26-28 Priority
	// 25 Reserved
	// 24 DGN High
	// 16-23 DGN (H)
	//8-15 DGN Low
	//0-7 SA
	var sa uint8 = 0x14
	var dgn uint32 = DGN_ADDRESS_CLAIMED
	// add eff
	var dgnWithEFF uint32 = (dgn|constants.CAN_EFF_FLAG)<<8 | uint32(sa)

	// Set the MSBytes
	binary.LittleEndian.PutUint32(raw[0:], dgnWithEFF)
	var length uint8 = 8
	raw[4] = length
	raw[5] = 0
	raw[6] = 0
	raw[7] = 0

	// b0 serial # lsb -- Optional. Required if multiple nodes from the same manufacturer may be present on the network.
	// b1 serial #
	// b2(0-4) serial # msb
	// b2(5-7) mfgr code LSB -- Required. Code obtained from SAE or RVIA.
	// b3  mfgr code
	// b4 (0-2) node instance -- For devices implementing multiple RV-C nodes (normally 0)
	// b4 (3-7) function instance -- Optional, intended to allow multiple instances of the same RV-C node, normally 0
	// b5 compatability field -- Optional, normally 0
	// b6 (0) reserved ZERO -- Always 0
	// b6 (1-7) compatability field -- Optional, normally 0
	// b7 (0-3) Optional, normally 0
	// b7 (4-6) Always 0
	// b7 (7) Required. 1 – Node supports address claiming 0 – Node uses a fixed source address

	raw[8] = 0x11         // b0 snr
	raw[9] = 0x11         // b1 snr
	raw[10] = 0xBF        // b2 snr(0-4) mfgr lsb(5-7)
	raw[11] = 0x01        // b3 mfgr code
	raw[12] = 0b0101_0101 // b4 node(0-2) function(3-7)
	raw[13] = 0b0101_1010 // b5 compatability field 5A
	raw[14] = 0xF9        // b6 reserved(0) compat2 (1-7)  1111 1001 (f9) 0111 1100 7c
	// F9
	// 1111 1001
	//  111 1100

	raw[15] = 0xA5

	f.BuildCanFrameX()
	claimed.Init(&f)

	//if raw[15] == 0 && claimed.SourceAddress == 0 {
	//	t.Errorf("WTF\n")
	//}

	{
		var expectedSNR uint32 = 0x01F1111
		if claimed.serialNumber != expectedSNR {
			t.Errorf(" address claimed SNR dos not match expected  %x got %x\n", expectedSNR, claimed.serialNumber)
		}
		fmt.Printf("Addr Clmd: %s\n", claimed.String())
	}
	{
		var expectedMFGRCode uint32 = 0x0d
		if claimed.mfgrCode != expectedMFGRCode {
			t.Errorf(" mfgr code SNR dos not match expected  %x got %x\n", expectedMFGRCode, claimed.mfgrCode)
		}
	}
	{
		var expectedNode uint8 = 0x05
		if claimed.nodeInstance != expectedNode {
			t.Errorf(" instance dos not match expected  %x got %x\n", expectedNode, claimed.mfgrCode)
		}
	}
	{
		var expectedFunction uint8 = 0x0A
		if claimed.functionInstance != expectedFunction {
			t.Errorf(" function dos not match expected  %x got %x\n", expectedFunction, claimed.functionInstance)
		}
	}
	{
		var expectedCompatabilityField1 uint8 = 0x5a
		if claimed.compatabilityField1 != expectedCompatabilityField1 {
			t.Errorf(" compatabilityField1 dos not match expected  %x got %x\n", expectedCompatabilityField1, claimed.compatabilityField1)
		}
	}
	{
		var expectedreserved uint8 = 0x1
		if claimed.reserved1 != expectedreserved {
			t.Errorf(" compatabilityField1 dos not match expected  %x got %x\n", expectedreserved, claimed.reserved1)
		}
	}
	{
		var compat2 uint8 = 0x7c
		if claimed.compatabilityField2 != compat2 {
			t.Errorf(" compatabilityField2 dos not match expected  %x got %x\n", compat2, claimed.compatabilityField2)
		}
	}

	// A5 1010 0101 -- 101
	// // b7 (0-3) Optional, normally 0 (5))
	//	// b7 (4-6) Always 0 (2)
	//	// b7 (7) Required. 1 – Node supports address claiming 0 – Node uses a fixed source address (1)
	{
		var compat3 uint8 = 5
		if claimed.compatabilityField3 != compat3 {
			t.Errorf(" compatabilityField2 dos not match expected  %x got %x\n", compat3, claimed.compatabilityField3)
		}
	}
	{
		var compat4 uint8 = 2
		if claimed.compatabilityField4 != compat4 {
			t.Errorf(" compatabilityField2 dos not match expected  %x got %x\n", compat4, claimed.compatabilityField4)
		}
	}
	{
		var cap uint8 = 1
		if claimed.arbitraryAddressCapable != cap {
			t.Errorf(" arbitraryAddressCapable dos not match expected  %x got %x\n", cap, claimed.arbitraryAddressCapable)

		}
	}

	fmt.Printf("Addr Clmd: %s\n", claimed.String())
}

func TestAddressClaimedCreateFrame(t *testing.T) {

	var addrClaimed = AddressClaimed{}
	addrClaimed.DGN = DGN_ADDRESS_CLAIMED
	addrClaimed.Name = DGNName(addrClaimed.DGN)
	addrClaimed.Priority = 0x06
	addrClaimed.SourceAddress = 0xFE
	addrClaimed.SetSourceAddress(0x11)
	// set AFTER setting up the DGN so it can be update correctly
	addrClaimed.SetSerialNumber(0xFF)
	addrClaimed.SetManufacturerCode(0x22)
	addrClaimed.SetNodeInstance(3)
	addrClaimed.SetFunctionInstance(0x03)
	addrClaimed.SetCompatabilityField1(0xAA)
	addrClaimed.SetCompatabilityField2(0x4A)
	addrClaimed.SetCompatabilityField3(0x09)
	addrClaimed.SetCompatabilityField4(0x05)
	var createdFrame = addrClaimed.CreateFrame()

	var dataBytesFrame = RvcFrame{}
	var bytes = dataBytesFrame.GetMessage()
	for i := 0; i < len(bytes); i++ {
		dataBytesFrame.GetMessage()[i] = createdFrame.GetMessage()[i]
	}

	dataBytesFrame.BuildCanFrameX()
	var rebuiltAddressClaimed = AddressClaimed{}
	rebuiltAddressClaimed.Init(&dataBytesFrame)

	fmt.Printf("OUT = %x\n", createdFrame.MessageBytes)
	fmt.Printf("OUT = %x\n", dataBytesFrame.MessageBytes)

	if !addrClaimed.Equals(&rebuiltAddressClaimed) {
		t.Errorf("Initial and reconstructed address claim commands dont match %s != %s\n", addrClaimed.String(), rebuiltAddressClaimed.String())
	}

}
