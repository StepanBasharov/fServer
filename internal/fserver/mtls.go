package fserver

import (
	"fmt"
	"fullserver/internal/tools"
	"log/slog"
	"net/http"
	"sync"
)

type MTLsServer struct {
	Server *http.ServeMux
	Port   int
}

func (server *MTLsServer) Run(group *sync.WaitGroup) error {
	defer group.Done()
	slog.Info("mtls server starting")
	err := http.ListenAndServe(tools.GetPort(server.Port), server.Server)
	if err != nil {
		return fmt.Errorf("can't create mtls server")
	}
	return nil
}
