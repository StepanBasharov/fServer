package fServer

import (
	"fullserver/internal/fconfig"
	"fullserver/internal/fserver"
	"net/http"
	"sync"
)

type FServer struct {
	HTTP   fserver.HTTPServer
	HTTPS  fserver.HTTPSServer
	MTLs   fserver.MTLsServer
	Config fconfig.FServerConfig
}

func NewFServer(cfg fconfig.FServerConfig) *FServer {
	var fServer FServer
	if cfg.HTTP.UseHTTP {
		fServer.HTTP = fserver.HTTPServer{Server: http.NewServeMux(), Port: cfg.HTTP.HTTPPort}
	}
	if cfg.HTTPS.UseHTTPS {
		fServer.HTTPS = fserver.HTTPSServer{Server: http.NewServeMux(), Port: cfg.HTTPS.HTTPSPort}
	}
	if cfg.MTLs.UseMTLs {
		fServer.MTLs = fserver.MTLsServer{Server: http.NewServeMux(), Port: cfg.MTLs.MTLsPort}
	}
	fServer.Config = cfg

	return &fServer
}

func (fs *FServer) Run() {
	var wg sync.WaitGroup
	if fs.Config.HTTP.UseHTTP {
		wg.Add(1)
		go fs.HTTP.Run(&wg)
	}
	if fs.Config.HTTPS.UseHTTPS {
		wg.Add(1)
		go fs.HTTPS.Run(&wg)
	}
	if fs.Config.MTLs.UseMTLs {
		wg.Add(1)
		go fs.MTLs.Run(&wg)
	}
	wg.Wait()
}
