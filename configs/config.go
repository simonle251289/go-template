package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"template/utils"
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
	JWT struct {
		AccessTTL  int `yaml:"accessTTL"`
		RefreshTTL int `yaml:"refreshTTL"`
	}
	RSAPrivateKey        string
	RSAPublicKey         string
	RSARefreshPrivateKey string
	RSARefreshPublicKey  string
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
	config.RSAPrivateKey = utils.LoadFileString("/configs/rsa.pem")
	config.RSAPublicKey = utils.LoadFileString("/configs/rsa.pub")
	config.RSARefreshPrivateKey = utils.LoadFileString("/configs/rsa-refresh.pem")
	config.RSARefreshPublicKey = utils.LoadFileString("/configs/rsa-refresh.pub")
	return config
}
