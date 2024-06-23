package main

import "fmt"

type ServerOption func(*Server)

type Server struct {
	Host    string
	Port    int
	Timeout int
}

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}
func WithTimeOut(timeout int) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{ // 这三个是默认值
		Host:    "localhost",
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range options {
		opt(server)
	}
	return server
}

func optionsModuleTest() {
	srv := NewServer(WithHost("baidu.com"), WithPort(8081), WithTimeOut(60))
	fmt.Println("srv:", srv)
}
