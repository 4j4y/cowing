package main

import (
	"os"
	"github.com/4j4y/cowing/cowin"
	"github.com/4j4y/cowing/notify"
)

var BLACKLIST = []int{ 582783 }

func main() {
	queryType := os.Args[1]
	identifierID := os.Args[2]
	frequency := os.Args[3]
	songToPlay := "vaccinenew.mp3"
	daysToSearch := ""
	if len(os.Args) > 4 {
		daysToSearch = os.Args[4]
	}

	if len(os.Args) > 5 {
		songToPlay = os.Args[5]
	}

	ch := make(chan cowin.Response)
	go notify.Subscribe(ch)
	cowin.InitRecursiveFetch(queryType, identifierID, frequency, daysToSearch,songToPlay, BLACKLIST, ch)
}
