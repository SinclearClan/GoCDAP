package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/tidwall/jsonc"
)

type Config struct {
	Calendar struct {
		Type     string `json:"type"`
		Url      string `json:"url"`
		Path     string `json:"path"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"calendar"`
	Discord struct {
		AppId string `json:"appId"`
	} `json:"discord"`
}

func parseConfigFromJson(filename string) (config *Config) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	config = new(Config)
	configJson := jsonc.ToJSON([]byte(configFile))

	err = json.Unmarshal(configJson, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func Init() *Config {
	cfg := parseConfigFromJson("./config.jsonc")
	return cfg
}
