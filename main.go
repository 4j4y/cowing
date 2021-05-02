package main

import (
	"github.com/4j4y/cowing/cowin"
	"github.com/4j4y/cowing/notify"
	"os"
)

var BLACKLIST = []int{ 582783 }

func main() {
	queryType := os.Args[1]
	identifierID := os.Args[2]
	frequency := os.Args[3]
	daysToSearch := ""
	if len(os.Args) > 4 {
		daysToSearch = os.Args[4]
	}
	ch := make(chan cowin.SessionResponse)
	go notify.Subscribe(ch)
	cowin.InitRecursiveFetch(queryType, identifierID, frequency, daysToSearch, BLACKLIST, ch)
}
