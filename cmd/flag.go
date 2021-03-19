package main 

import (
	"os"
)

func flag() (string, string, error) {

	managerID := os.Getenv("MANAGER_ID")

	leagueID := os.Getenv("LEAGUE_ID")

	return managerID, leagueID, nil
}