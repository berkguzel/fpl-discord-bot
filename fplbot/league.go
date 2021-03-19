package fplbot

import (
	"fmt"
)
var (
	message string
)

func Standings(leagueID string) (string, error) {

	c := Authentication()
	standings, err := c.GetStandings(leagueID)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range standings {
		message = message + fmt.Sprintf("%d) %s, %d \n", k+1, v.EntryName, v.Total)
	}

	return message, nil
}

func Team(name string) (string, error) {

	a := Authentication()

	c := Client{}
	fmt.Println(c.LeagueID)
	team, err := a.GetStandings(c.LeagueID)
	if err != nil {
		return "Could not fetch standings", err
	}

	for k, v := range team {
		if name == v.EntryName {
			message = message + fmt.Sprintf("%s, \n%d, \n%d, \n%d, \n%d, \n%d", v.EntryName, k+1,  v.ID, v.Rank, v.Total, v.EventTotal)
		}
	}

	return message, nil
}