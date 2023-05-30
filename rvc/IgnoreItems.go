package rvc

// There are some DGN/instances where we don't know what they are and they create a lot of distracting noise
// for the time being - we're going to filter them out. We want to come back around to this later and don't want
// this filter list to get lost inside some other file...

//var IgnoreItems map[interface{}]bool
//
//func init() {
//	IgnoreItems[DGNInstanceKey{DGN_DC_DIMMER_STATUS_3, 59}] = true
//}

func Ignore(itemIF *RvcItemIF) bool {
	if (*itemIF).GetDGN() == DGN_DC_DIMMER_STATUS_3 && (*itemIF).GetInstance() == 59 {
		return true
	} else {
		return false
	}
}
