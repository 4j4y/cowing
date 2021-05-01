package main

import (
	"encoding/json"
	"fmt"
	"github.com/gen2brain/beeep"
	"io/ioutil"
	"net/http"
	"time"
)

type CovidData struct {
	Centers []struct {
		CenterID     int    `json:"center_id"`
		Name         string `json:"name"`
		StateName    string `json:"state_name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Lat          int    `json:"lat"`
		Long         int    `json:"long"`
		From         string `json:"from"`
		To           string `json:"to"`
		FeeType      string `json:"fee_type"`
		Sessions     []struct {
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
		VaccineFees []struct {
			Vaccine string `json:"vaccine"`
			Fee     string `json:"fee"`
		} `json:"vaccine_fees,omitempty"`
	} `json:"centers"`
}

func main() {
	usingPin := false
	usingDistrictID := true
	pin := "332001"
	did := "513"
	loc, _ := time.LoadLocation("Asia/Calcutta")
	fmt.Print("\033[H\033[2J")
	for usingPin {
		todayString := GetDate(loc)
		callCowinUsingPin(pin, todayString)
		time.Sleep(15 * time.Minute)
		fmt.Print("\033[H\033[2J")
	}
	for usingDistrictID {
		todayString := GetDate(loc)
		callCowinUsingDid(did, todayString)
		time.Sleep(2 * time.Minute)
		fmt.Print("\033[H\033[2J")
	}

}

func GetDate(loc *time.Location) string {
	istNow := time.Now().In(loc)
	fmt.Printf("Script last pinged at %v\n", istNow)
	year, month, day := istNow.Date()
	todayString := fmt.Sprintf("%02d-%02d-%d", day, month, year)
	return todayString
}

func callCowinUsingPin(pin string, date string) {
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=" + pin + "&date=" + date
	callCowin(url)
}

func callCowinUsingDid(did string, date string) {
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=" + did + "&date=" + date
	callCowin(url)
}

func callCowin(url string) {
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("accept", "application/json, text/plain, */*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var covidData CovidData
	err = json.Unmarshal(body, &covidData)
	if err != nil {
		fmt.Println("-----------")
		fmt.Println(err)
		fmt.Println("-----------")
	}
	totalCenters := len(covidData.Centers)
	for i := 0; i < totalCenters; i++ {
		totalSessions := len(covidData.Centers[i].Sessions)
		for j := 0; j < totalSessions; j++ {
			minAgeLimit := covidData.Centers[i].Sessions[j].MinAgeLimit
			if minAgeLimit == 18 {
				if covidData.Centers[i].CenterID != 582783 && // blacklist a center
					covidData.Centers[i].Sessions[j].AvailableCapacity > 0 { // remove unusable centers
					fmt.Println("+++++++++++++++Center Information+++++++++++++++++")
					fmt.Printf("Center ID:\t\t %d \n", covidData.Centers[i].CenterID)
					fmt.Printf("Center Name:\t\t %s \n", covidData.Centers[i].Name)
					fmt.Printf("Center Pincode:\t\t %d \n", covidData.Centers[i].Pincode)
					fmt.Printf("Center Lat:\t\t %d \n", covidData.Centers[i].Lat)
					fmt.Printf("Center Long:\t\t %d \n", covidData.Centers[i].Long)
					fmt.Printf("Date:\t\t\t %s \n", covidData.Centers[i].Sessions[j].Date)
					fmt.Printf("Available Capacity:\t %d \n", covidData.Centers[i].Sessions[j].AvailableCapacity)
					fmt.Printf("Vaccine type:\t %s \n", covidData.Centers[i].Sessions[j].Vaccine)
					msgBody := fmt.Sprintf("Center Name: %s \nAvailable Capacity: %d", covidData.Centers[i].Name, covidData.Centers[i].Sessions[j].AvailableCapacity)
					err := beeep.Alert("Found a center", msgBody, "assets/information.png")
					if err != nil {
						panic(err)
					}
				}
			}

		}
	}
}
