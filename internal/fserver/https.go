package fserver

import (
	"fmt"
	"fullserver/internal/tools"
	"log/slog"
	"net/http"
	"sync"
)

type HTTPSServer struct {
	Server *http.ServeMux
	Port   int
}

func (server HTTPSServer) Run(group *sync.WaitGroup) error {
	defer group.Done()
	slog.Info("https server starting")
	err := http.ListenAndServe(tools.GetPort(server.Port), server.Server)
	if err != nil {
		return fmt.Errorf("can't create https server")
	}
	return nil
}
