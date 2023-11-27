package config

import (
	"fmt"
)

type System struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Env             string `yaml:"env"`
	LoginLimitTime  int    `yaml:"login-limit-time"`
	LoginLimitCount int    `yaml:"login-limit-count"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
