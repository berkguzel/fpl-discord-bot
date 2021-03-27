package main

import (
	"os"
)

func flag() (string, string, error) {

	token := os.Getenv("TOKEN")
	leagueID := os.Getenv(("LEAGUE_ID"))

	return token, leagueID, nil
}
