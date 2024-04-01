package test

import (
	"github.com/stellaraf/go-utils/environment"
)

var Env *EnvType = nil

type EnvType struct {
	PanoramaHost     string `env:"PANORAMA_HOST"`
	FirewallIP       string `env:"FIREWALL_IP"`
	PanoramaUsername string `env:"PANORAMA_USERNAME"`
	PanoramaPassword string `env:"PANORAMA_PASSWORD"`
	FirewallUsername string `env:"FIREWALL_USERNAME"`
	FirewallPassword string `env:"FIREWALL_PASSWORD"`
	HostnamePattern  string `env:"HOSTNAME_PATTERN"`
}

func LoadEnv() (*EnvType, error) {
	var env EnvType
	err := environment.Load(&env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func init() {
	env, err := LoadEnv()
	if err != nil {
		panic(err)
	}
	Env = env
}
