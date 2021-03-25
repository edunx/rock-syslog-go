package syslog

import (
	"github.com/edunx/lua"
	tp "github.com/edunx/rock-transport-go"
)

type Server struct {
	lua.Super

	protocol  string
	listen    string
	format    string
	transport []tp.Tunnel
}
