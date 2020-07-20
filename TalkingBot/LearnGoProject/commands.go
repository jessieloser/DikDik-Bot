package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

func OnHelp(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//build string to display help/info/commands
	embed := buildEmbed(config.CommandTitle, config.Commands)
	s.ChannelMessageSendEmbed(msg.ChannelID, &embed)

}

func OnJokeThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	fmt.Println(arg[1])
	//confirm channelid exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(jokes))
		s.ChannelMessageSend(arg[1], jokes[randomIndex])
	} else {
		s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
	}

}

func OnJokeHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(jokes))
	s.ChannelMessageSend(msg.ChannelID, jokes[randomIndex])
}
func OnFactsThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channelid exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(facts))
		s.ChannelMessageSend(arg[1], facts[randomIndex])
	} else {
		s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
	}

}

func OnFactsHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(facts))
	s.ChannelMessageSend(msg.ChannelID, facts[randomIndex])
}

func OnSet(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channel id is in list and print id
	if len(arg[:]) > 1 {
		s.ChannelMessageSend(arg[1], arg[2])
	} else {
		s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
	}
}
