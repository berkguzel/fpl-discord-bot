package fplbot

import ( 
	
	"github.com/berkguzel/fpl-go/fpl"

)

func Authentication() *fpl.Client {
	c := fpl.NewClient(nil)

	return c
}