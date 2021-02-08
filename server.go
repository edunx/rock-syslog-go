package syslog

import (
	"encoding/json"
	"errors"
	"fmt"
	pub "github.com/edunx/rock-public-go"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func (s *Server) Start() error {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)

	switch s.protocol {
	case "tcp":
		server.ListenTCP(s.listen)
	case "udp":
		server.ListenUDP(s.listen)
	case "tcp/udp":
		server.ListenUDP(s.listen)
		server.ListenTCP(s.listen)
	default:
		return errors.New("invalid protocol , must be tcp , udp , tcp/udp; got " + s.protocol)
	}

	server.Boot()
	for logParts := range channel {
		switch  s.format {
		case "json":
			if v, e := json.Marshal(logParts); e == nil {
				s.transport.Push( v )
			} else {
				pub.Out.Err("syslog-go err: %v" , e)
			}

		default:
			s.transport.Push( fmt.Sprintf("%v\n" , logParts ))
		}
	}

	return nil
}

func (s *Server) Close() {

}

func (s *Server) Reload() {

}
