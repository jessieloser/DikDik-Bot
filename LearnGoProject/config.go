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
type channel struct{
	Topic	string
}

const configPath = "./dikdik-config.yml"

func GetConfig() DikDikConfig {
	var prefix = "/"
	defaultConfig := DikDikConfig{
		Prefix:       prefix,
		CommandTitle: "Commands",
		Commands: []string{
			prefix + "+say MentionChannel message to send to channel",
			"*Activate message sending to MentionedChannel. All messages you send hereafter will be send to this channel*",
			prefix + "-say",
			"*Deactivate message sending to MentionChannel*",
			prefix + "delete",
			"*Delete last sent message while say is active*",
			prefix + "jokeHere",
			"*Post a joke in current channel*",
			prefix + "jokeThere MentionChannel",
			"*Send joke to the MentionedChannel*",
			prefix + "factsHere",
			"*Post facts in current channel*",
			prefix + "factsThere MentionChannel",
			"*Send facts to the MentionedChannel*",
			prefix + "help",
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
