package server

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func New(host string, port int, option ...Option) *Server {
	svr := &Server{
		host: host,
		port: port,
	}

	// 可変長引数をfor文で回し、関数を実行。引数にsvr を渡すことでsvr を更新する
	for _, opt := range option {
		opt(svr)
	}

	return svr
}

func (s Server) Start() error {
	return nil
}

// server を引数に取る関数に名前をつける
type Option func(server *Server)

// Option を返すようにして、引数のserver を更新する関数を返す
func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
