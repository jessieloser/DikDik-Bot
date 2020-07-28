package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func OnMessage(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	//confirm prefix is correct
	if len(msg.GuildID) == 0 || !strings.HasPrefix(msg.Content, config.Prefix) {
		return
	}
	//split string
	args := strings.Split(msg.Content, " ")
	if len(args[:]) > 1 {
		if args[0] == config.Prefix+"+say" {
			if len(args[:]) > 2 {
				args[2] = strings.Join(args[2:], " ")
				//makes sure all spaces are trimmed from front and back
				for i, arg := range args {

					args[i] = strings.TrimSpace(arg)
					//stops trimming so it doesnt remove spaces from arg[2]
					if args[i] == args[2] {
						break
					}
				}
			}
		}

		if len(args[:]) > 1 {
			if len(args[1]) > 19 && strings.Contains(args[1], "<") {
				runes2 := []rune(args[1])
				args[1] = string(runes2[2:20])
				_, err := s.State.Channel(args[1])
				if err != nil {
					fmt.Println(args[1], err)
					// Could not find channel.
					s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
					return
				}
			} else
			{
				_, err := s.State.Channel(args[1])
				if err != nil {
					fmt.Println(args[1], err)
					// Could not find channel.
					s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
					return
				}
			}
		}
	}

	//cut the / off the first string
	runes := []rune(args[0])
	args[0] = string(runes[1:])
	//check first arg to decide how to proceed
	switch args[0] {
	case "+say":
		OnSet(s, msg, args[:])
		break
	case "-say":
		OnUnset(s, msg, args[:])
	case "jokeThere":
		OnJokeThere(s, msg, args[:])
		break
	case "factsThere":
		OnFactsThere(s, msg, args[:])
		break
	case "factsHere":
		OnFactsHere(s, msg)
		break
	case "jokeHere":
		OnJokeHere(s, msg)
		break
	default:
		OnHelp(s, msg)
	}
}

func OnReady(s *discordgo.Session, ready *discordgo.Ready) {
	//confirms bot is ready
	testing := fmt.Sprintf("ready your %s\n", ready.User.Username)
	fmt.Printf(testing)
}
