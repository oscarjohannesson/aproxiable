package aproxiable

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	version string = "v1.0.0"
)

type Aproxiable struct {
	CertificatePath    string
	CertificateKeyPath string

	ListeningAddr string
	router        *mux.Router
	middlewares   []mux.MiddlewareFunc
	reverseProxys []*ReverseProxy
}

type AproxiableOptionFunc func(*Aproxiable) error

//NewAproxiable creates a new Aproxiable reverse-proxy
func NewAproxiable(options ...AproxiableOptionFunc) (*Aproxiable, error) {
	a := &Aproxiable{
		router: mux.NewRouter(),
	}

	//@TODO add create from functions
	return a, nil
}

//NewAproxiableFromConfig creates a Aproxiable from a Aproxiable struct @todo fill struct with missing, to make a standard configuration
func NewAproxiableFromConfig(config *Config) (*Aproxiable, error) {
	a := Aproxiable{
		router: mux.NewRouter(),
	}

	err := a.loadConfig(config)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

//NewAproxiableFromConfigPath creates a new Aproxiable reverse-proxy from yaml file
func NewAproxiableFromConfigPath(configPath string) (*Aproxiable, error) {

	config, err := parseConfigFromPath(configPath)

	a, err := NewAproxiableFromConfig(config)
	if err != nil {
		return nil, err
	}

	return a, err
}

//Start starts the Aproxiable reverse proxy, locks goroutine
func (a *Aproxiable) Start() {
	if len(a.reverseProxys) == 0 {
		panic("will not start proxy without any proxys configured")
	}

	if a.router == nil {
		panic("Cannot start the reverse proxy, need a vaild muxrouter the handle incomming request")
	}

	a.addReverseProxysToRouter(a.reverseProxys...)
	a.addMiddlewareToRouter(a.middlewares...)
	if a.CertificatePath == "" && a.CertificateKeyPath == "" {
		//add proper logging
		fmt.Printf("no certificate specified, starting http server on %s\n", a.ListeningAddr)
		http.ListenAndServe(a.ListeningAddr, a.router)

		return
	}
	//certificate detected, serve https
	fmt.Printf("certificate detected starting htpps server on %s\n", a.ListeningAddr)
	http.ListenAndServeTLS(a.ListeningAddr, a.CertificatePath, a.CertificateKeyPath, a.router)

}
