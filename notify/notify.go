package notify

import (
	"fmt"
	"github.com/4j4y/cowing/cowin"
	"github.com/gen2brain/beeep"
)

func Subscribe(ch chan cowin.SessionResponse)  {
	for {
		session := <-ch
		printIt(&session)
	}
}

func printIt(sessionData *cowin.SessionResponse) {
	fmt.Println("+++++++++++++++Center Information+++++++++++++++++")
	fmt.Printf("Center ID:\t\t %d \n", sessionData.Center.CenterID)
	fmt.Printf("Center Name:\t\t %s \n", sessionData.Center.Name)
	fmt.Printf("Center Pincode:\t\t %d \n", sessionData.Center.Pincode)
	fmt.Printf("Center Lat:\t\t %d \n", sessionData.Center.Lat)
	fmt.Printf("Center Long:\t\t %d \n", sessionData.Center.Long)
	fmt.Printf("Date:\t\t\t %s \n", sessionData.Session.Date)
	fmt.Printf("Available Capacity:\t %f \n", sessionData.Session.AvailableCapacity)
	fmt.Printf("Vaccine type:\t %s \n", sessionData.Session.Vaccine)
	msgBody := fmt.Sprintf("Center Name: %s \nAvailable Capacity: %f", sessionData.Center.Name, sessionData.Session.AvailableCapacity)
	err := beeep.Alert("Found a center", msgBody, "assets/information.png")
	if err != nil {
		panic(err)
	}
}

