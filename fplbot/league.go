package fplbot

import (
	"fmt"
)

func Standings(leagueID string) (string, error) {

	c := Authentication()
	var message string
	standings, err := c.GetStandings(leagueID)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range standings {
		message = message + fmt.Sprintf("%d) %s, %d \n", k+1, v.EntryName, v.Total)
	}

	return message, nil
}
