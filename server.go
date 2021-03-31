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
					s.write( v )
				} else {
					pub.Out.Err("syslog-go err: %v" , e)
				}
			case "line":
				s.write( fmt.Sprintf("%v" , item ))
			}
		}
	}(channel)

	s.obj = server

	return nil
}

func (s *Server) write( v interface{} ) {
	n := len(s.IO)
	for i := 0; i< n ;i++ {
		if err := s.IO[i].Write( v ); err != nil {
			pub.Out.Err("%s write io fail , err: %v" , s.name , err )
		}
	}
}

func (s *Server) Close() {
	pub.Out.Err("%s stop succeed" , s.name)
	s.obj.Kill()
}

func (s *Server) Type() string {
	return "syslog.server"
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Proxy( info string , v interface{}) {
}