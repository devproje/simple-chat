package config

import (
	"encoding/json"
	"github.com/devproje/plog/log"
	"os"
)

type Config struct {
	Logging    bool   `json:"logging"`
	ServerName string `json:"server_name"`
	Database   struct {
		Clusters   []string `json:"clusters"`
		KeySpace   string   `json:"keyspace"`
		Credential struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"credential"`
	} `json:"database"`
}

const filename = "config.json"
const sample = `{
	"logging": false,
	"server_name": "Simple Chat",
	"database": {
		"clusters": [],
		"keyspace": "",
		"credential": {
			"username": "",
			"password": ""
		}
	}
}`

func Load() *Config {
	raw, err := os.ReadFile(filename)
	if err != nil {
		file, err := os.Create(filename)
		if err != nil {
			log.Errorln(err)
			return nil
		}

		_, err = file.Write([]byte(sample))
		if err != nil {
			log.Errorln(err)
			return nil
		}

		log.Warnln("config.json is created, please type configuration to file.")
		os.Exit(0)
	}

	var data Config
	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &data
}
