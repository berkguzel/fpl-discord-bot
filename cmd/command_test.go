package main

import (
	"testing"
)

func TestParseMessage(t *testing.T) {

	var test = []string{
		"! -m *****", "!-m-m",
		"!!!!!!!!!!!!!!!!!!!!!!!!!!", "!-m!!!",
	}

	for _, v := range test {

		got := parseMessage(v)
		if got != "" {
			t.Errorf("Could not parse message")
		}
	}
}

func TestParseCommand(t *testing.T) {

	var test = []string{
		"!helppp", "!thsiweek",
		"!standdings", "!standssss",
		"!11111111111111111111",
	}

	for _, v := range test {
		got := parseCommand(v)
		if got != "" {
			t.Errorf("Please enter a valid command")
		}

	}
}
