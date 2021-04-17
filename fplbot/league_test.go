package fplbot

import (
	"testing"
)

func TestStandings(t *testing.T) {

	var test = []string{
		"aa", "", "---",
		"***", "    .",
	}

	for _, v := range test {

		got, _ := Standings(v)
		if got != "" {
			t.Errorf("Please enter a valid league id")
		}

		gottw, _ := ThisWeek(v)
		if gottw != "" {
			t.Errorf("Please enter a valid league id")
		}
	}
}

func TestThisWeek(t *testing.T) {

}
