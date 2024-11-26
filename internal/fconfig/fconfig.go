package fconfig

type FServerConfig struct {
	HTTP struct {
		UseHTTP  bool
		HTTPPort int
	}
	HTTPS struct {
		UseHTTPS     bool
		HTTPSPort    int
		CertFilePath string
		KeyFilePath  string
	}
	MTLs struct {
		UseMTLs        bool
		MTLsPort       int
		CaCertFilePath string
	}
}

func NewConfig(useHTTP, useHTTPS, useMTLs bool, HTTPPort, HTTPSPort, MTLsPort int) FServerConfig {
	fConfig := FServerConfig{}
	if useHTTP {
		fConfig.HTTP = struct {
			UseHTTP  bool
			HTTPPort int
		}{UseHTTP: useHTTP, HTTPPort: HTTPSPort}
	}
	if useHTTPS {
		fConfig.HTTPS = struct {
			UseHTTPS     bool
			HTTPSPort    int
			CertFilePath string
			KeyFilePath  string
		}{UseHTTPS: useHTTPS, HTTPSPort: HTTPSPort, CertFilePath: "", KeyFilePath: ""}
	}
	if useMTLs {
		fConfig.MTLs = struct {
			UseMTLs        bool
			MTLsPort       int
			CaCertFilePath string
		}{UseMTLs: useMTLs, MTLsPort: MTLsPort, CaCertFilePath: ""}
	}
	return fConfig
}
