package fplbot

import (
	"fmt"

)
var (
	message string
)

func Standings(leagueID string) (string){

	c := Authentication()
	standings, err := c.GetStandings(leagueID)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range standings {
		message = message + fmt.Sprintf("%d) %s, %d \n", k+1, v.EntryName, v.Total)
	}

	return message
}