package notify

import (
	"fmt"
	"github.com/4j4y/cowing/cowin"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
	"github.com/faiface/beep/mp3"
	"github.com/gen2brain/beeep"
)

func Subscribe(ch chan cowin.Response)  {
	for {
		session := <-ch
		printIt(&session.SessionResponse, session.SongToPlay)
	}
}
func playSong(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	start := time.Now()
	speaker.Play(streamer)

	elapsed := time.Since(start)

	for elapsed.Seconds() < 60 {
		elapsed = time.Since(start)
	}
}

func printIt(sessionData *cowin.SessionResponse, songToPlay string) {
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
	playSong(songToPlay)
}

