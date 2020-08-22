package internal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Variables used for command line parameters
type Bot struct {
	config    DikDikConfig
	botPrefix string
	client    *discordgo.Session
	allVars   Variables1
}

type Variables1 struct {
	//holds the author name and id of channel they are writing to
	m map[string]string
	//holds the author name and id of last message sent
	dm map[string]string
	//holds the time since last edit
	tm map[string]time.Time
	//holds the author and the id of the channel they are writing in
	cm map[string]string

	ourBool bool

	//duration between last message and current message
	between time.Duration
	//time until timeout and deactivate say automatically
	sayoffTime float64
	//create joke array
	jokelist []string
	//create fact array
	factlist []string
}

func Start(config DikDikConfig) {

	//initialize client
	client, _ := discordgo.New("Bot " + config.Token)
	//set variables
	varbs := Variables1{
		ourBool:    false,
		sayoffTime: 5,
	}

	//create bot
	bot := Bot{
		config:    config,
		botPrefix: config.Prefix,
		client:    client,
		allVars:   varbs,
	}

	//loads files
	bot.allVars.jokelist = readFile(config.CSVPathJokes)
	bot.allVars.factlist = readFile(config.CSVPathFacts)
	fmt.Println("files read")

	//confirm client opened properly
	if err := client.Open(); err != nil {
		log.Fatalln("Failed to connect to Discord. Is token correct?\n" + err.Error())
	}
	//create message map
	bot.allVars.m = make(map[string]string)
	//create prior message map
	bot.allVars.dm = make(map[string]string)
	//create current channel map
	bot.allVars.cm = make(map[string]string)
	//create timestamp
	bot.allVars.tm = make(map[string]time.Time)

	//confirms bot is ready
	fmt.Println("ready your dikdik")

	//initialize handlers
	client.AddHandler(bot.onMessage)
	client.AddHandler(bot.onEdit)
}

//embed for the help menu thing
func (bot Bot) buildEmbed() discordgo.MessageEmbed {
	embed := discordgo.MessageEmbed{}
	//check to see if embed is already built
	if bot.allVars.ourBool == false {
		embed.Color = 0x1385ef
		embed.Title = "Commands"
		//only join when first creating
		embed.Description =  "`/+say channelName [message to send to channel]`\n"+
		 "Activate message sending to MentionedChannel. All messages you send hereafter will be send to this channel\n"+
		 "`/-say`\n"+
		 "Deactivate message sending to MentionChannel\n"+
		 "`/delete`\n"+
		 "Delete last sent message while say is active\n"+
		 "`/jokeHere`\n"+
		 "Post a joke in current channel\n"+
		 "`/jokeThere MentionChannel`\n"+
		 "Send joke to the MentionedChannel\n"+
		 "`/factsHere`\n"+
		 "Post facts in current channel\n"+
		 "`/factsThere MentionChannel`\n"+
		 "Send facts to the MentionedChannel\n"+
		 "`/status`\n"+
		 "Confirm if say is currently active\n"+
		 "`/help`"
		bot.allVars.ourBool = true
		return embed
	}
	return embed
}

func readFile(filename string) []string {
	//read file
	file, err := ioutil.ReadFile(filename)
	//check for errors
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	//create string to hold file contents
	var str = string(file)
	//split string on comma
	var txtlines = strings.Split(str, "|")
	//return array
	return txtlines[:]
}
