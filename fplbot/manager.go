package fplbot

import (
	"fmt"
	"strconv"
)

func ManagerID(leagueID string, teamName string) (int, error) {

	c := Authentication()
	manager, err := c.GetTeamInfoInLeague(leagueID, teamName)
	if err != nil {
		return 0, fmt.Errorf("Could not find manager id \n %s", err)
	}

	return manager.Entry, nil
}

func Manager(leagueID string, teamName string) (string, error) {

	c := Authentication()
	var message string

	id, err := ManagerID(leagueID, teamName)
	if err != nil {
		fmt.Println(err)
	}
	m, err := c.GetManager(strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
	}

	message = message + fmt.Sprintf(" Manager: %s\n Overall Rank: %d\n Overall Points: %d\n Event Points: %d",
		m.PlayerFirstName, m.SummaryOverallRank, m.SummaryOverallPoints, m.SummaryEventPoints)

	return message, nil
}
