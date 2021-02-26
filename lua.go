package syslog

import (
	"github.com/edunx/lua"
)

func LuaInjectApi(L *lua.LState, parent *lua.LTable) {
	sysTab := L.CreateTable(0, 2)
	LuaInjectServerApi(L , sysTab)

	L.SetField(parent, "syslog", sysTab)
}

