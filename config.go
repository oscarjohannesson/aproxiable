package aproxiable

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Config contains everything needed for setting up
type Config struct {
	//version (Not really needed but looks pretty cool)
	AproxiableVersion  string          `yaml:"aproxiableVersion"`
	ListeningAddr      string          `yaml:"listeningAddr"`
	CertificatePath    string          `yaml:"certificatePath"`
	CertificateKeyPath string          `yaml:"certificateKey"`
	ReverseProxys      []*ReverseProxy `yaml:"reverse-proxys"`
}

type ReverseProxy struct {
	Name          string `yaml:"name"`
	ListeningPath string `yaml:"listeningPath"` //if you want a seperate listing addr for your
	TargetAddr    string `yaml:"targetAddr"`    //including path
}

//Middleware @TODO have standard middlewares defined and a list with middleware functions that can be actived inert for now
type Middleware struct {
	Name string `yaml:"name"`
}

//ParseConfig parses an aproxiable yaml configuration
func parseConfigFromPath(configPath string) (*Config, error) {

	filename, _ := filepath.Abs(configPath)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

//loadConfig loads the reverse proxy configuration from a config, overwrites latest configuration, changes during runtime is not supportet
func (a *Aproxiable) loadConfig(config *Config) error {
	if config == nil {
		return fmt.Errorf("config cannot be empty")
	}
	//@TODO add a default proxy configuration

	a.configToAproxiable(config)
	return nil
}

func (a *Aproxiable) configToAproxiable(config *Config) error {

	a.ListeningAddr = config.ListeningAddr
	a.CertificatePath = config.CertificatePath
	a.CertificateKeyPath = config.CertificateKeyPath

	//adding reverse proxies
	for _, reverseProxy := range config.ReverseProxys {
		fmt.Println()
		a.reverseProxys = append(a.reverseProxys, reverseProxy)
	}

	return nil
}
