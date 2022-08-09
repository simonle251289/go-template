package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Debug    bool `yaml:"debug"`
	Port     int  `yaml:"port"`
	Database struct {
		Postgres struct {
			Driver   string `yaml:"driver"`
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Database string `yaml:"database"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		}
	}
}

func LoadConfig() Config {
	wd, _ := os.Getwd()
	file, err := ioutil.ReadFile(wd + "/configs/template.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
