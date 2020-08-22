# DikDik-Bot
DikDik bot is a Discord bot that says something in another channel. Tell it what to say in channel A and it will send it to channel B. Once say is active all messages will go to channel B until it is deactivated. The bot will automatically deactivate if there is no messages from the sending user for more then 5 minutes.

## Setup
Download [Go](https://go-lang.org/).

        $ cd ./cmd/DikDik-Bot
        $ go build
        $ ./DikDik-Bot
        # Setup dikdik-config.yml
        
        
## Config File
 In dikdik-config.yml
 
         bot_token: <your bot token>
         bot_prefix: <prefix>
         csv_path_jokes: <path where Jokes.txt is located>
         csv_path_facts: <path where Facts.txt is located>
        
## Usage
Available commands
        
Command | Description
------------ | -------------
/+say channelName [message to send to channel] | Activate message sending to MentionedChannel. All messages you send hereafter will be send to this channel
/-say | Deactivate message sending to MentionChannel
/delete|Delete last sent message while say is active
/jokeHere|Post a joke in current channel
/jokeThere MentionChannel|Send joke to the MentionedChannel
/factsHere|Post facts in current channel
/factsThere MentionChannel|Send facts to the MentionedChannel
/status| Confirm if say is currently active
/help| Command list
