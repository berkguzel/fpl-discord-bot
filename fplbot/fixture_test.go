package fplbot

import (
	"testing"
)

func TestFixture(t *testing.T) {

	got := GetFixture(1)
	if got == "" {
		t.Errorf("Could do not find the fixture")
	}

}
