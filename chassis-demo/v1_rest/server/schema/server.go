package schema

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

// RespData rest Response data structure
type RespData struct {
}

type Server struct {
}

func (s *Server) SayHello(ref *restful.Context) {
	println("say hello...")
	err := ref.Write([]byte("say hello..."))
	if err != nil {
		return
	}
}

// URLPatterns is a fixed method and only structs that implement it can be registered as a schemas for the rest service
func (s *Server) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/sayhello", ResourceFunc: s.SayHello},
	}
}
