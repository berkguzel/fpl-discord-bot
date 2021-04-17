package main

import (
	"fmt"
	"strconv"
	"strings"

	fpl "github.com/berkguzel/fpl-discord-bot/fplbot"
	"github.com/bwmarrin/discordgo"
)

func parseMessage(message string) string {

	var name string
	var msg string
	parse := strings.Fields(string(message[3:]))

	_, leagueID, err := flag()
	if err != nil {
		return ""
	}

	command := parse[0]
	if command == "m" {

		if len(parse) > 2 {
			name = parse[1]
			for i := 2; i < len(parse); i++ {

				name = name + " " + parse[i]

			}
		} else {
			name = parse[1]

		}

		msg, err := fpl.Manager(leagueID, name)
		if err != nil {
			fmt.Println(err)
		}

		return msg
	}

	if command == "f" {
		gameWeek, _ := strconv.Atoi(parse[1])
		msg = fpl.GetFixture(gameWeek)

		return msg
	}

	return ""

}

// TODO detect empty strings
func parseCommand(message string) string {

	_, leagueID, err := flag()
	if err != nil {
		return ""
	}

	if message == "!standings" {

		msg, err := fpl.Standings(leagueID)
		if err != nil {
			return ""
		}

		return msg
	}

	if message == "!thisweek" {

		msg, err := fpl.ThisWeek(leagueID)
		if err != nil {
			return ""
		}

		return msg
	}

	if message == "!help" {

		msg := help()
		return msg
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
