package syslog

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
	tp "github.com/edunx/rock-transport-go"
)

func CheckTransports( v lua.LValue ) []tp.Tunnel {

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

	rc := make([]tp.Tunnel , size)
	for i:= 1 ; i<=size ; i++ {
		item := tab.RawGetInt(i)
		if item.Type() != lua.LTUserData {
			pub.Out.Err("id: %d  not transport userdata" , i)
			return nil
		}

		tun := tp.CheckTunnel( item.(*lua.LUserData) )
		rc[i - 1] = tun
	}

	return rc
}