package syslog

import pub "github.com/edunx/rock-public-go"

type Server struct {
	protocol  string
	listen    string
	format    string
	transport []pub.Transport
}
