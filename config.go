package syslog

import (
	"github.com/edunx/lua"
	"gopkg.in/mcuadros/go-syslog.v2"
)

type Server struct {
	lua.Super

	name      string
	protocol  string
	listen    string
	format    string
	IO        []lua.IO

	obj       *syslog.Server
}
