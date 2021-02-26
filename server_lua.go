package syslog

import (
	"github.com/edunx/lua"
)

const (
	SERVERMT string = "ROCK_SYSLOG_SERVER_GO_MT"
	//CLIENTMT string = "ROCK_SYSLOG_CLIENT_GO_MT"
)

func LuaInjectServerApi(L *lua.LState, parent *lua.LTable) {
	mt := L.NewTypeMetatable(SERVERMT)
	L.SetField(mt, "__index", L.NewFunction(serverGet))
	L.SetField(mt, "__newindex", L.NewFunction(serverSet))

	L.SetField(parent , "server", L.NewFunction(CreateServerUserdata))
}

func CreateServerUserdata(L *lua.LState) int {
	opt := L.CheckTable(1)

	s := &Server{
		protocol: opt.CheckString("protocol", "udp"),
		listen:   opt.CheckString("listen", "0.0.0.0:514"),
		format:   opt.CheckString("format" , "json"),
		transport: CheckTransports(opt.RawGetString("transport")),
	}

	if e := s.Start(); e != nil {
		L.RaiseError("start syslog server fail , err:%v", e)
		return 0
	}

	ud := L.NewUserDataByInterface(s , SERVERMT)
	L.Push(ud)
	return 1
}

func serverGet(L *lua.LState) int {
	return 0
}

func serverSet(L *lua.LState) int {
	return 0
}

func (s *Server) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(s, SERVERMT)
}


