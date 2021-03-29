package fplbot

import (
	"testing"
)

func TestManagerID (t *testing.T) {

	var test = []string {
		"aa", "", "---", 
		"***", "    .",
		"???", "  -*  ",
	}

	for _, v := range test {

		got, _ := ManagerID(v, v) 
		if got != 0 {
			t.Errorf("Could not find the manager")
		}

		gotManager, _ := Manager(v,v)
		if  gotManager != "" {
			t.Errorf("Could not find the manager")
		}
	}
}
