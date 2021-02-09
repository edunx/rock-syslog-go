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
	go func(channel syslog.LogPartsChannel){
		for item := range channel {
			switch s.format {
			case "json":
				if v, e := json.Marshal( item ); e == nil {
					s.Push( v )
				} else {
					pub.Out.Err("syslog-go err: %v" , e)
				}
			case "line":
				s.Push( fmt.Sprintf("%v" , item ))
			}
		}
	}(channel)

	return nil
}

func (s *Server) Push( v interface{} ) {
	for _ , tp := range	s.transport {
		tp.Push( v )
	}
}

func (s *Server) Close() {

}

func (s *Server) Reload() {

}
