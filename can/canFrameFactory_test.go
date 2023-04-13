package can

import (
	"rvctomqtt/constants"
	"testing"
)

func TestCanFrameFactory(t *testing.T) {
	var fac = CanFrameFactory{}
	var fr = fac.Create()
	var messageLength = len((*(*fr).GetMessage()))

	if constants.MAX_MESSAGE != int32(messageLength) {

		t.Errorf("Can message error, data length expected = %d got %d ", constants.MAX_MESSAGE, messageLength)

	}

}
