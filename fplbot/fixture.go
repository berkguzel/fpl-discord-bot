package fplbot

import (
	"fmt"

	"github.com/enescakir/emoji"
)

var (
	home        string
	away        string
	kickoffTime string
)

func GetTeam(home int, away int) (string, string) {

	codeOfTeam := map[int]string{
		1:  "Arsenal",
		2:  "Aston Villa",
		3:  "Brentford",
		4:  "Brighton and Hove Albion",
		5:  "Burnley",
		6:  "Chelsea",
		7:  "Crystal Palace",
		8:  "Everton",
		9:  "Leicestser United",
		10: "Leeds United",
		11: "Liverpool",
		12: "Manchester City",
		13: "Manchester United",
		14: "Newcastle United",
		15: "Norwich City",
		16: "Southampton",
		17: "Totthenham Hotspur",
		18: "Watford",
		19: "West Ham United",
		20: "Wolverhampton Wanderers",
	}

	return codeOfTeam[home], codeOfTeam[away]

}
func GetFixture(gameWeek int) string {

	c := Authentication()

	fixture, err := c.GetFixture()
	if err != nil {
		fmt.Println(err)
	}

	var message string

	for _, v := range fixture[0] {
		if v.Event == float64(gameWeek) {
			home, away = GetTeam(v.TeamH, v.TeamA)
			kickoffTime = fmt.Sprintf("%v", v.KickoffTime)
			message = message + "\n" + string(emoji.SoccerBall) + home + " " + "-" + " " + away + "  " +
				string(emoji.AlarmClock) + " " + kickoffTime

		}
	}
	return message
}
