package fplbot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/enescakir/emoji"
)

func Standings(leagueID string) (string, error) {

	c := Authentication()
	var message string

	standings, err := c.GetStandings(leagueID)

	if err != nil {
		return "", errors.New("Could not fetch standings")
	}

	for k, v := range standings {
		message = message + fmt.Sprintf("%d) %s, %d \n", k+1, v.EntryName, v.Total)
	}
	message = string(emoji.SoccerBall) + string(emoji.SoccerBall) + string(emoji.SoccerBall) + string(emoji.SoccerBall) + "\n" + message

	return message, nil
}

func ThisWeek(leagueID string) (string, error) {

	c := Authentication()
	var message string
	var leader string

	standings, err := c.GetStandings(leagueID)
	if err != nil {
		return "", errors.New("Could not fetch standings")
	}

	count := 0

	for _, v := range standings {

		id, err := ManagerID(leagueID, v.EntryName)
		if err != nil {
			return "", errors.New("Please enter a valid team name")
		}

		m, err := c.GetManager(strconv.Itoa(id))
		if err != nil {
			return "", errors.New("Could not find the manager")
		}

		if count == 0 {
			message = fmt.Sprintf("%s Game Week: %d %s \n", emoji.SoccerBall, m.CurrentEvent, emoji.SoccerBall)
		}

		if m.SummaryEventPoints > count {
			count = m.SummaryEventPoints
			leader = v.EntryName
		}
		
		message = message + fmt.Sprintf("%s	%d \n", v.EntryName, m.SummaryEventPoints)

	}

	message = fmt.Sprintf("%s %s %s is on fire %s %s \n", emoji.Fire, emoji.Fire, leader, emoji.Fire, emoji.Fire) + message

	return message, nil

}
