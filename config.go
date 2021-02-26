package syslog

import (
	tp "github.com/edunx/rock-transport-go"
)

type Server struct {
	protocol  string
	listen    string
	format    string
	transport []tp.Tunnel
}
