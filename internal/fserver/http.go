package fserver

import (
	"fmt"
	"fullserver/internal/tools"
	"log/slog"
	"net/http"
	"sync"
)

type HTTPServer struct {
	Server *http.ServeMux
	Port   int
}

func (server *HTTPServer) Run(group *sync.WaitGroup) error {
	defer group.Done()
	slog.Info("http server starting")
	err := http.ListenAndServe(tools.GetPort(server.Port), server.Server)
	if err != nil {
		return fmt.Errorf("can't create http server")
	}
	return nil
}
