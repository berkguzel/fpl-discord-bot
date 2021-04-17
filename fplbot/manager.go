package fplbot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/enescakir/emoji"
)

func ManagerID(leagueID string, teamName string) (int, error) {

	c := Authentication()
	manager, err := c.GetTeamInfoInLeague(leagueID, teamName)
	if err != nil {
		return 0, errors.New("Could not find manager id")
	}

	return manager.Entry, nil
}

func Manager(leagueID string, teamName string) (string, error) {

	c := Authentication()
	var message string

	id, err := ManagerID(leagueID, teamName)
	if err != nil {
		return "", errors.New("Could not find the manager id")
	}

	m, err := c.GetManager(strconv.Itoa(id))
	if err != nil {
		return "", errors.New("Could not find the manager")
	}

	message = message + fmt.Sprintf("%s Manager: %s\n Overall Rank: %d\n Overall Points: %d\n Event Points: %d",
		emoji.Man, m.PlayerFirstName, m.SummaryOverallRank, m.SummaryOverallPoints, m.SummaryEventPoints)

	return message, nil
}
