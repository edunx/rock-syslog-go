package syslog

import (
	"github.com/edunx/lua"
)

func LuaInjectApi(L *lua.LState, parent *lua.UserKV) {
	syslog := &lua.UserKV{}
	syslog.Set("server" , lua.NewGFunction( createSyslogServer ))
	parent.Set("syslog", syslog )
}