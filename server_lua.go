package syslog

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

func (self *Server) ToLightUserData(L *lua.LState) *lua.LightUserData {
	return L.NewLightUserData( self )
}

func (self *Server) CheckIO( L *lua.LState , v lua.LValue ) {
	if v.Type() != lua.LTTable {
		L.RaiseError("invalid type , must be IO table , got %s" , v.Type().String() )
		return
	}

	tab := v.(*lua.LTable)
	n := tab.Len()
	if n == 0 {
		pub.Out.Err("not found IO")
		return
	}
	self.IO = make([]lua.IO , n )

	for i:= 1 ; i<= n ; i++ {
		val := tab.RawGetInt(i)
		if val.Type() != lua.LTLightUserData {
			self.IO = nil
			L.RaiseError("id: %d  not IO userdata" , i)
			return
		}
		self.IO[i - 1] = val.(*lua.LightUserData).CheckIO(L)
	}
}

func createSyslogServer (L *lua.LState , args *lua.Args ) lua.LValue {
	opt := args.CheckTable(L , 1)
	s := &Server{
		protocol: opt.CheckString("protocol", "udp"),
		listen:   opt.CheckString("listen", "0.0.0.0:514"),
		format:   opt.CheckString("format" , "json"),
		name:     opt.CheckString("name" , "syslog.server"),
	}

	s.CheckIO(L , opt.RawGetString("IO") )

	if e := s.Start(); e != nil {
		L.RaiseError("start syslog server fail , err:%v", e)
		return lua.LNil
	}

	return s.ToLightUserData( L )
}