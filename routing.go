package aproxiable

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func addReverseProxy(proxies ...*ReverseProxy) AproxiableOptionFunc {
	return func(a *Aproxiable) error {
		for _, proxy := range proxies {
			a.reverseProxys = append(a.reverseProxys, proxy)
		}
		return nil
	}
}

//newReverseProxy creates a new reverse proxy
func (a *Aproxiable) addReverseProxysToRouter(proxyConfigs ...*ReverseProxy) error {

	for _, proxyConfig := range proxyConfigs {
		fmt.Printf("setting up routing for %v\n", proxyConfig)
		remote, err := url.Parse(proxyConfig.TargetAddr)
		if err != nil {
			return err
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		//@TODO add more custom configurations for the proxy in the configuration
		//Define the director func
		//This is a good place to log, for example or change the incomming request
		/* 	proxy.Director = func(req *http.Request) {
			fmt.Println(req)
		}
		*/
		//adding custom transport layer (using default as standard) @TODO define extra transport stuff in the config to add here later
		/* 		proxy.Transport = http.DefaultTransport
		 */

		a.router.HandleFunc(proxyConfig.ListeningPath, ProxyRequestHandler(proxy)) /* .Name(proxyConfig.Name) */

	}

	return nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("\n request:%+v", r.Header)
		fmt.Printf("\n response:%+v", w)
		proxy.ServeHTTP(w, r)
	}
}
