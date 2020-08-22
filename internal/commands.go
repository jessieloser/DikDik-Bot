package internal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"time"
)

//help command
func (bot Bot) onHelp(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//build string to display help/info/commands
	embed := bot.buildEmbed()
	_, err := s.ChannelMessageSendEmbed(msg.ChannelID, &embed)
	if err != nil {
		fmt.Println(err)
	}
}

//jokeThere command
func (bot Bot) onJokeThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channel ID exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(bot.allVars.jokelist))
		_, err := s.ChannelMessageSend(arg[1], bot.allVars.jokelist[randomIndex])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err := s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
		if err != nil {
			fmt.Println(err)
		}
	}
}

//jokeHere command
func (bot Bot) onJokeHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(bot.allVars.jokelist))
	_, err := s.ChannelMessageSend(msg.ChannelID, bot.allVars.jokelist[randomIndex])
	if err != nil {
		fmt.Println(err)
	}
}

//factsThere command
func (bot Bot) onFactsThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channel ID exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(bot.allVars.factlist))
		_, err := s.ChannelMessageSend(arg[1], bot.allVars.factlist[randomIndex])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err := s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
		if err != nil {
			fmt.Println(err)
		}
	}
}

//factsHere command
func (bot Bot) onFactsHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(bot.allVars.factlist))
	_, err := s.ChannelMessageSend(msg.ChannelID, bot.allVars.factlist[randomIndex])
	if err != nil {
		fmt.Println(err)
	}

}

//+say command
func (bot Bot) onSet(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirms there is a channel in the message
	if len(arg[:]) > 1 {
		//create map record of author and channel id where message is sent
		bot.allVars.m[msg.Author.Username] = arg[1]
		//create map record of author and channel id where message is from
		bot.allVars.cm[msg.Author.Username] = msg.ChannelID
		//set timestamp to last message sent
		bot.allVars.tm[msg.Author.Username] = time.Now()
		//confirm there is something to post after setting the say
		if len(arg[:]) > 2 {
			//post in other channel
			message, err := s.ChannelMessageSend(arg[1], arg[2])
			if err != nil {
				fmt.Println(err)
			}
			//record message id that was posted to other channel
			bot.allVars.dm[bot.allVars.m[msg.Author.Username]] = message.ID
			_, errd := s.ChannelMessageSend(msg.ChannelID, "You are now sending messages to <#"+arg[1]+">")
			if errd != nil {
				fmt.Println(errd)
			}

		} else {
			_, err := s.ChannelMessageSend(msg.ChannelID, "You are currently sending messages to <#"+bot.allVars.m[msg.Author.Username]+">")
			if err != nil {
				fmt.Println(err)
			}

		}
	} else {
		_, err := s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
		if err != nil {
			fmt.Println(err)
		}

	}
}

//text while say command is active
func (bot Bot) onText(s *discordgo.Session, msg *discordgo.MessageCreate) {

	if _, exists := bot.allVars.tm[msg.Author.Username]; exists {
		//sets current time
		currentTime := time.Now()
		//sets old time based on info in map
		oldtime := bot.allVars.tm[msg.Author.Username]
		//finds the difference between the two times
		diff := currentTime.Sub(oldtime)
		//confirms it hasnt been more then x# of minutes since last message
		if diff.Minutes() < bot.allVars.sayoffTime {
			//check if there is an attachment
			if len(msg.Attachments) == 0 {
				//check if message contains content
				if msg.Content != "" {
					//posts all messages to other channels
					message, err := s.ChannelMessageSend(bot.allVars.m[msg.Author.Username], msg.Content)
					if err != nil {
						fmt.Println(err)
					}
					//record message id that was posted to other channel
					bot.allVars.dm[bot.allVars.m[msg.Author.Username]] = message.ID
				}
			} else {
				//msg contains an attachment
				bot.onAttach(s, msg.Attachments[0], msg)
			}
		} else {
			//if its been longer then x# of minutes delete records
			_, err := s.ChannelMessageSend(msg.ChannelID,
				" Say automatically turned off for user "+msg.Author.Username+" after "+
					strconv.FormatFloat(bot.allVars.sayoffTime, 'f', 0, 64)+" minutes")
			if err != nil {
				fmt.Println(err)
			}

			delete(bot.allVars.dm, bot.allVars.m[msg.Author.Username])
			delete(bot.allVars.m, msg.Author.Username)
			delete(bot.allVars.tm, msg.Author.Username)
			delete(bot.allVars.cm, msg.Author.Username)
			return
		}
	} else {
		return
	}
}

//-say command
func (bot Bot) onUnset(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//clear the channel topic
	//turns out there is a long cool down on this so I cant use this
	//s.ChannelEditComplex(msg.ChannelID, &discordgo.ChannelEdit{
	//	Topic: "",
	//})
	if err, exists := bot.allVars.m[msg.Author.Username]; exists {
		if err != "" {
			fmt.Println(err)
		}
		//if user exists delete record in map
		_, err := s.ChannelMessageSend(msg.ChannelID,
			"You are no longer sending messages to channel <#"+bot.allVars.m[msg.Author.Username]+">")
		if err != nil {
			fmt.Println(err)
		}

		//clear all maps of user data
		delete(bot.allVars.dm, bot.allVars.m[msg.Author.Username])
		delete(bot.allVars.m, msg.Author.Username)
		delete(bot.allVars.tm, msg.Author.Username)
		delete(bot.allVars.cm, msg.Author.Username)
	} else {
		//if user doesnt exist return
		_, err := s.ChannelMessageSend(msg.ChannelID, "Say is not currently active for "+msg.Author.Username)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//deletes message last posted in channel
func (bot Bot) onDelete(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//checks if user exists- last message contains a value
	if _, exists := bot.allVars.dm[bot.allVars.m[msg.Author.Username]]; exists {
		s.ChannelMessageDelete(bot.allVars.m[msg.Author.Username], bot.allVars.dm[bot.allVars.m[msg.Author.Username]])
		s.ChannelMessageSend(msg.ChannelID, "The message has been deleted")
	} else {
		//if user doesnt exist return
		s.ChannelMessageSend(msg.ChannelID, "There is no prior message available to be deleted")
	}
}

//checks if say is active
func (bot Bot) onStatus(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//if user exists in map say active
	if err, exists := bot.allVars.m[msg.Author.Username]; exists {
		if err != "" {
			fmt.Println(err)
		}
		_, err := s.ChannelMessageSend(msg.ChannelID, "Say is currently active for "+msg.Author.Username+" in channel <#"+
			bot.allVars.m[msg.Author.Username]+">")
		if err != nil {
			fmt.Println(err)
		}
		_, errd := s.ChannelMessageSend(msg.ChannelID, "Thanks for checking in. I'm still a piece of garbage")
		if errd != nil {
			fmt.Println(errd)
		}

	} else {
		//user doesnt exist in map- not active
		_, err := s.ChannelMessageSend(msg.ChannelID, "Say is not currently active for "+msg.Author.Username)
		if err != nil {
			fmt.Println(err)
		}
		_, errd := s.ChannelMessageSend(msg.ChannelID, "Thanks for checking in. I'm still a piece of garbage")
		if errd != nil {
			fmt.Println(errd)
		}

	}
}

func (bot Bot) onAttach(s *discordgo.Session, attmsg *discordgo.MessageAttachment, msg *discordgo.MessageCreate) {
	//checks to see if attachment message contains text/a title
	if msg.Content != "" {
		//posts message content and url to other channel
		message, err := s.ChannelMessageSend(bot.allVars.m[msg.Author.Username], msg.Content+" "+attmsg.URL)
		if err != nil {
			fmt.Println(err)
		}
		//record message id that was posted to other channel
		bot.allVars.dm[bot.allVars.m[msg.Author.Username]] = message.ID
	} else {
		//message doesnt contain content
		message, err := s.ChannelMessageSend(bot.allVars.m[msg.Author.Username], attmsg.URL)
		if err != nil {
			fmt.Println(err)
		}
		bot.allVars.dm[bot.allVars.m[msg.Author.Username]] = message.ID
	}
}
