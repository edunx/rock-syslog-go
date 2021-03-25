package syslog

import (
	"github.com/edunx/lua"
)

func (self *Server) ToLightUserData(L *lua.LState) *lua.LightUserData {
	return L.NewLightUserData( self )
}

func createSyslogServer (L *lua.LState , args *lua.Args ) lua.LValue {
	opt := args.CheckTable(L , 1)
	s := &Server{
		protocol: opt.CheckString("protocol", "udp"),
		listen:   opt.CheckString("listen", "0.0.0.0:514"),
		format:   opt.CheckString("format" , "json"),
		transport: CheckTransports(opt.RawGetString("transport")),
	}

	if e := s.Start(); e != nil {
		L.RaiseError("start syslog server fail , err:%v", e)
		return lua.LNil
	}

	return s.ToLightUserData( L )
}