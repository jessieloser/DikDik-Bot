package internal

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"log"
)

type DikDikConfig struct {
	Token        string   `yaml:"bot_token"`
	Prefix       string   `yaml:"bot_prefix"`
	CSVPathJokes string   `yaml:"csv_path_jokes"`
	CSVPathFacts string   `yaml:"csv_path_facts"`
}

const ConfigPath = "dikdik-config.yml"

func GetConfig() (config DikDikConfig) {
	config = DikDikConfig{
		Prefix: "/",
	}

	err := util.GetConfig(ConfigPath, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config
}
