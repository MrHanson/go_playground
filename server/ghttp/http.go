package ghttp

import (
	"fmt"
	"net/http"
)

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

func (s *sdkHttpServer) Start(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}

func NewSdkHttpServer(name string) *sdkHttpServer {
	return &sdkHttpServer{
		Name: name,
	}
}

func HttpRun() {
	s := NewSdkHttpServer("demo-server")

	s.Route("/", sayHello)

	s.Start("8000", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello world")
	w.Write([]byte(string("hello world")))
}
