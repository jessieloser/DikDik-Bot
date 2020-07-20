package main

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"log"
)

type DikDikConfig struct {
	Token        string            `yaml:"bot_token"`
	Prefix       string            `yaml:"bot_prefix"`
	CommandTitle string            `yaml:"command_title"`
	Commands     []string          `yaml:"commands"`
	Channels     map[string]string `yaml:"stored_channels"`
}

const configPath = "./dikdik-config.yml"

func GetConfig() DikDikConfig {
	var prefix = "/"
	defaultConfig := DikDikConfig{
		Prefix:       prefix,
		CommandTitle: "Commands",
		Commands: []string{
			prefix + "set channelName message to send to channel",
			prefix + "jokeHere",
			prefix + "jokeThere channelName",
			prefix + "factsHere",
			prefix + "factsThere channelName",
			prefix + "help DikDik",
		},
	}

	err := util.GetConfig(configPath, &defaultConfig)

	if err != nil {
		log.Fatalln(err)
	}

	return defaultConfig
}

func (config DikDikConfig) Save() {
	util.Save(configPath, &config)
}
