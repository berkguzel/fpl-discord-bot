package fplbot

import (
	"fmt"
)

func GetTeam(code int) string {

	codeOfTeam := map[int]string{
		1: "Arsenal F.C.",
		2: "Aston Villa F.C.",
		3: "Brighton and Hove Albion F.C.",
		4: "Burnley",
		5: "Chelsea",
		6: "Crystal Palace",
		7: "Everton",
		8: "Fulham",
		9: "Leeds United",
		10: "Leicestser United",
		11: "Liverpool",
		12: "Manchester City", 
		13: "Manchester United",
		14: "Newcastle United",
		15: "Sheffield United",
		16: "Southampton",
		17: "Totthenham Hotspur",
		18: "West Bromwich Albion",
		19: "West Ham United",
		20: "Wolverhampton Wanderers",

	}

	return codeOfTeam[code]

	
}
func GetFixture() {

	c := Authentication()

	fixture, err := c.GetFixture()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range fixture {
		fmt.Println(v[0].Event, v[0].KickoffTime, v[0].TeamA, v[0].TeamH)
		fmt.Println(v[1].Event, v[1].KickoffTime, v[1].TeamA, v[1].TeamH)
		fmt.Println(v[30].Event, v[30].KickoffTime, v[30].TeamA, v[30].TeamH)
	}
}