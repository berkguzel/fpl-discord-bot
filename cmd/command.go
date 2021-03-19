package main

import (
	"github.com/bwmarrin/discordgo"
	fpl "github.com/berkguzel/fpl-discord-bot/fplbot"
)

var (
	message string
	
)

func parseMessage(args string) string{

	command := string(args[3])
	message = string(args[5:])

	if command == "t" {
		msg, err := fpl.Team(message, )
		if err != nil {
			return ""
		}
		return msg
	}

	return ""

}
// TODO detect empty strings
func parseCommand(args string) (string) {


	if args == "!standings" {
		msg, err := fpl.Standings("1859418")
		if err != nil {
			return ""
		}

		return msg
	}

	if string(args[0:9]) == "!leagueID" {

		leagueID := string(args[9:])
		var c fpl.Client
		c.AuthFPL(leagueID, "")
	}

	if string(args[0:10]) == "!managerID" {

		var manager map[string]string
		managerID := string(args[10:])
		var c fpl.Client
		c.AuthFPL("", managerID)
		manager["first"] = managerID
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
