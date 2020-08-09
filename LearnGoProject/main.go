package main

import (
	util "github.com/Floor-Gang/utilpkg"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var config DikDikConfig
var m map[string]string
var dm map[string]string
var bool = false

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

//embed for the help menu thing
func buildEmbed(s string, cmd []string) discordgo.MessageEmbed {
	//check to see if embed is already built
	if bool == false {
		embed := discordgo.MessageEmbed{}
		embed.Color = 0x1385ef
		embed.Title = s
		//only join when first creating
		cmd[0] = strings.Join(cmd[0:], " \n")
		embed.Description = cmd[0]
		bool = true
		return embed
	} else {
		embed := discordgo.MessageEmbed{}
		embed.Color = 0x1385ef
		embed.Title = s
		embed.Description = cmd[0]
		return embed

	}
}
