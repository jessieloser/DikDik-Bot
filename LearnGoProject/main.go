package main

import (
	util "github.com/Floor-Gang/utilpkg"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

var config DikDikConfig
//holds the author name and id of channel they are writing to
var m map[string]string
//holds the author name and id of last message sent
var dm map[string]string
//holds the time since last edit
var tm map[string]time.Time
//holds the author and the id of the channel they are writing in
var cm map[string]string
var bool = false
//duration between last message and current message
var between time.Duration
//time until timeout and deactivate say automatically
var sayoffTime float64 = 5

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
