package server

import (
	"log"
	"time"
)

type Server struct {
	param serverParam
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	return nil
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

// serverParam はserver package の中でだけ使用するのでexport しなくて良い
func NewBuilder(host string, port int) *serverParam {
	return &serverParam{
		host: host,
		port: port,
	}
}

// serverParam が持つメソッドになっていて、返り値の型がserverParam なので別のメソッドを呼び出せる
func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}

func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

func (sb *serverParam) Build() *Server {
	svr := &Server{param: *sb}
	return svr
}
