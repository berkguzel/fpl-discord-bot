package main

func help() string {

	message :=
		`!standings					: lists standings in league
!help							  : lists help message		
!thisweek					  : lists points of the game week
! -f {game week}	     : lists the fixture of the week
! -m {team name}	   : lists manager's informations`

	return message
}
