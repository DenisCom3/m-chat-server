package auth

import (
	"fmt"
	"net"
)

type Auth struct {
	Host string `yaml:"host" env-required:"true"`
	Port int64 `yaml:"port" env-required:"true"`
}

func (a Auth) Address() string {
	return net.JoinHostPort(a.Host, fmt.Sprintf("%v", a.Port))
}