package syslog

import (
	"github.com/edunx/lua"
)

const (
	SERVERMT string = "ROCK_SYSLOG_SERVER_GO_MT"
	CLIENTMT string = "ROCK_SYSLOG_CLIENT_GO_MT"
)

func LuaInjectApi(L *lua.LState, parent *lua.LTable) {
	server := L.NewTypeMetatable(SERVERMT)
	L.SetField(server, "__index", L.NewFunction(serverGet))
	L.SetField(server, "__newindex", L.NewFunction(serverSet))

	client := L.NewTypeMetatable(CLIENTMT)
	L.SetField(client, "__index", L.NewFunction(clientGet))
	L.SetField(client, "__newindex", L.NewFunction(clientSet))

	sysTab := L.CreateTable(0, 2)
	L.SetField(sysTab, "server", L.NewFunction(CreateSyslogServerUserdata))
	L.SetField(sysTab, "client", L.NewFunction(CreateSyslogClientUserdata))

	L.SetField(parent, "syslog", sysTab)
}

func CreateSyslogServerUserdata(L *lua.LState) int {
	opt := L.CheckTable(1)

	s := &Server{
		protocol: opt.CheckString("protocol", "udp"),
		listen:   opt.CheckString("listen", "0.0.0.0:514"),
		format:   opt.CheckString("format" , "json"),
		transport: CheckTransports( opt.RawGetString("transport")),
	}

	if e := s.Start(); e != nil {
		L.RaiseError("start syslog server fail , err:%v", e)
		return 0
	}

	ud := L.NewUserDataByInterface(s , SERVERMT)
	L.Push(ud)
	return 1
}

func CreateSyslogClientUserdata(L *lua.LState) int {
	return 0
}

func serverGet(L *lua.LState) int {
	return 0
}

func serverSet(L *lua.LState) int {
	return 0
}

func clientGet(L *lua.LState) int {
	return 0
}

func clientSet(L *lua.LState) int {
	return 0
}

func (s *Server) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(s, SERVERMT)
}
