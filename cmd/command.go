package main

import (
	"fmt"
	fpl "github.com/berkguzel/fpl-discord-bot/fplbot"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	message string
	id      string
	name    string
)

func parseMessage(message string) string {

	parse := strings.Fields(string(message[3:])) 

	command := parse[0]

	if command == "m" {

		name = parse[1]
		id := parse[0]

		msg, err := fpl.Manager(id, name)
		if err != nil {
			fmt.Println(err)
		}

		return msg
	}

	return ""

}

// TODO detect empty strings
func parseCommand(message string) string {

	_, leagueID, _ := flag() 
	if message == "!standings" {

		msg, err := fpl.Standings(leagueID)
		if err != nil {
			return ""
		}

		return msg
	}

	if message == "!help" {

		msg := help()

		return msg
	}

	if message == "!Yine ibneleşti" {
		return "galatasaray"
	}

	if message == "!Zaten hep ibneydi" {
		return "galatasaray"
	}


	return ""
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if string(m.Content[0]) == "!" && len(m.Content) > 4 {

		if string(m.Content[2]) == "-" {
			s.ChannelMessageSend(m.ChannelID, parseMessage(string(m.Content)))
		} else {
			s.ChannelMessageSend(m.ChannelID, parseCommand(string(m.Content)))
		}

	}

}
