package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"server"`
}

var once sync.Once
var instance *Config

func Init() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			log.Fatalln(err)
		}
	})
	return instance
}
