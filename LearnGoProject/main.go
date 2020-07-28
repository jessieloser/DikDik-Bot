package main

import (
	util "github.com/Floor-Gang/utilpkg"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var config DikDikConfig

func main() {
	config = GetConfig()
	client, _ := discordgo.New("Bot " + config.Token)

	client.AddHandler(OnMessage)
	client.AddHandler(OnReady)

	if err := client.Open(); err != nil {
		log.Fatalln("Failed to connect to Discord. Is token correct?")
	}
	util.KeepAlive()
}

func buildEmbed(s string, cmd []string) discordgo.MessageEmbed {
	embed := discordgo.MessageEmbed{}
	embed.Color = 0x1385ef
	embed.Title = s
	cmd[0] = strings.Join(cmd[0:], " \n")
	embed.Description = cmd[0]

	return embed
}
