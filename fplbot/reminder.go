package fplbot

import (
	"fmt"
	"time"
)

var (
	weekCounter int = 0
	currentEvent interface{}
	nextFixture time.Time
	game interface{}
	nextGame  time.Time
	homeT int 
	awayT int
	message string
	kickOffTimeStr string
)

func Reminder  () (time.Time, time.Time){
	c := Authentication()
	currentTime := time.Now().UTC()
	fixture, _ := c.GetFixture()
	for _, v := range fixture[0] {

		event, kickOffTime := v.Event, v.KickoffTime
		kickOffTimeStr := fmt.Sprintf("%v", kickOffTime)
		t, err := time.Parse(time.RFC3339, kickOffTimeStr)
		if err != nil {
			panic(err)
		}
		// find the current event(game week)
		if ( t.Sub(currentTime) > 0 && weekCounter == 0) {
			currentEvent = event
			weekCounter = weekCounter + 1
			//fmt.Println("current", currentTime.Day(), t.Day())
		
		} else if weekCounter == 1 && currentEvent != event{
			currentEvent = event
			game, homeT, awayT  = v.KickoffTime, v.TeamH, v.TeamA
			kickOffTimeStr = fmt.Sprintf("%v", game)
			nextGame, _ = time.Parse(time.RFC3339, kickOffTimeStr)
			weekCounter = 2

		} else if weekCounter == 2 && currentEvent != event {
			game = v.KickoffTime
			kickOffTimeStr = fmt.Sprintf("%v", game)
			nextFixture, _ = time.Parse(time.RFC3339, kickOffTimeStr)
			weekCounter = 0
			return nextGame, nextFixture
		}
		
	}
	
	return time.Time{}, time.Time{}
}

func RemindTimer() time.Time{
	nextGame, nextFixture = Reminder()
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println(err)
	}
	timeForMSK := nextGame.In(location)
	duration, _ := time.ParseDuration("-2.5h")
	remindTime := timeForMSK.Add(duration)
	
	return remindTime

}

func ReminderMessage() string{

	message = "Please check your team, there is only an hour left to pick team of this week"

	return message

}