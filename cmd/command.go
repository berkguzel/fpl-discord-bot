package main

import (

	"github.com/bwmarrin/discordgo"
	fpl "github.com/berkguzel/fpl-discord-bot/fplbot"
)


// TODO detect empty strings
func parseCommand(args string) (string) {

	if args == "!standings" {
		return fpl.Standings("")
	}

	return ""
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return 
	}
	
	if string(m.Content[0]) == "!" && len(m.Content) > 1{
		s.ChannelMessageSend(m.ChannelID, parseCommand(string(m.Content)))

	} 

	
}
