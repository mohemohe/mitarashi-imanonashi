package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	env struct {
		Mastodon struct {
			Server      string `required:"true"`
			AccessToken string `required:"true"`
		}
	}
)

var e *env

func GetEnv() *env {
	if e == nil {
		en := env{}
		if err := envconfig.Process("", &en); err != nil {
			panic(err)
		}
		e = &en
	}
	return e
}
