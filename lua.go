package syslog

import (
	"github.com/edunx/lua"
)

func LuaInjectApi(L *lua.LState, parent *lua.LTable) {
	syslog := &lua.UserKV{}

	syslog.Set("server" , lua.NewGFunction( createSyslogServer ))

	L.SetField(parent, "syslog", syslog )
}