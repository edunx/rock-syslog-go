package syslog

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

func CheckTransports( v lua.LValue ) []pub.Transport {

	if v.Type() != lua.LTTable {
		pub.Out.Err("not found transports")
		return nil
	}

	tab := v.(*lua.LTable)
	size := tab.Len()
	if size <= 0 {
		pub.Out.Err("not found transports")
		return nil
	}

	rc := make([]pub.Transport , size)
	for i:=1 ; i<=size ; i++ {
		item := tab.RawGetInt(i)
		if item.Type() != lua.LTUserData {
			pub.Out.Err("id: %d  not transport userdata" , i)
			return nil
		}

		tp := pub.CheckTransport( item.(*lua.LUserData) )
		rc[i - 1] = tp
	}

	return rc
}