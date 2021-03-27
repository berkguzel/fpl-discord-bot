package fplbot

import (
	"github.com/berkguzel/fpl-go/fpl"
)

type Client struct {
	ManagerID string
	LeagueID  string
}

func Authentication() *fpl.Client {

	x := fpl.NewClient(nil)

	return x
}

func (c *Client) AuthFPL(leagueID string, managerID string) *Client {

	c.LeagueID = leagueID
	c.ManagerID = managerID

	return c
}
