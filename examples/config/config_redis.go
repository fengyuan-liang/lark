package config

import "lark/pkg/conf"

type Config struct {
	Redis *conf.Redis
}

func GetConfig() (cfg *Config) {
	return
}
