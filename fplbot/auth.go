package fplbot

import (
	"github.com/berkguzel/fpl-go/fpl"
)

func Authentication() *fpl.Client {

	x := fpl.NewClient(nil)

	return x
}
