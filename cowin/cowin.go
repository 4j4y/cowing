package cowin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var BLACKLIST []int

type CovidData struct {
	Centers []Center `json:"centers"`
}
type Center struct {
	CenterID     int       `json:"center_id"`
	Name         string    `json:"name"`
	StateName    string    `json:"state_name"`
	DistrictName string    `json:"district_name"`
	BlockName    string    `json:"block_name"`
	Pincode      int       `json:"pincode"`
	Lat          int       `json:"lat"`
	Long         int       `json:"long"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	FeeType      string    `json:"fee_type"`
	Sessions     []Session `json:"sessions"`
	VaccineFees  []struct {
		Vaccine string `json:"vaccine"`
		Fee     string `json:"fee"`
	} `json:"vaccine_fees,omitempty"`
}

type Session struct {
	SessionID         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity float32  `json:"available_capacity"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}


type SessionResponse struct {
	Center  Center
	Session Session
}

type Response struct {
	SessionResponse SessionResponse
	SongToPlay string
}

func InitRecursiveFetch(queryType, identifierID, frequency, daysToSearch, songToPlay string, blacklist []int,
	ch chan Response) {
	BLACKLIST = blacklist
	var daySpan int
	var frequencyInMinutes time.Duration

	usingPin := false
	usingDistrictID := false
	pin := ""
	did := ""
	daySpan = 0
	frequencyInMinutes = 1

	if queryType == "pin" {
		usingPin = true
	} else if queryType == "did" {
		usingDistrictID = true
	} else {
		panic("provide either `pin` or `did`")
	}
	if len(identifierID) == 0 {
		panic("provide appropriate `pin` or `did`")
	} else if queryType == "pin" {
		pin = identifierID
	} else {
		did = identifierID
	}
	if len(frequency) > 0 {
		parseInt, err := strconv.ParseInt(frequency, 10, 64)
		if err != nil {
			panic("wrong format of frequency")
		}
		frequencyInMinutes = time.Duration(parseInt)
	}

	if len(daysToSearch) > 0 {
		days, err := strconv.Atoi(daysToSearch)
		if err != nil {
			panic("wrong format of day span to search")
		}
		daySpan = days
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")
	fmt.Print("\033[H\033[2J")

	for usingPin {

		for i := 0; i <= daySpan; i++ {
			fmt.Print("calling using Pin")
			sessionData := callCowinUsingPin(pin, GetDate(loc, i))
			if sessionData != nil {
				ch <- Response{
				  SessionResponse: *sessionData,
				  SongToPlay: songToPlay,
				}
			}
		}

		time.Sleep(frequencyInMinutes * time.Minute)
		fmt.Print("\033[H\033[2J")
	}
	for usingDistrictID {
		for i := 0; i <= daySpan; i++ {
			sessionData := callCowinUsingDid(did, GetDate(loc, i))
			if sessionData != nil {
				ch <- Response{
					SessionResponse: *sessionData,
					SongToPlay: songToPlay,
				}
			}
		}
		time.Sleep(frequencyInMinutes * time.Minute)
		fmt.Print("\033[H\033[2J")
	}
}

func GetDate(loc *time.Location, offset int) string {
	istNow := time.Now().AddDate(0, 0, 1*offset).In(loc)
	fmt.Printf("Script last pinged at %v\n", istNow)
	year, month, day := istNow.Date()
	todayString := fmt.Sprintf("%02d-%02d-%d", day, month, year)
	return todayString
}

func callCowinUsingPin(pin string, date string) *SessionResponse {
	fmt.Printf("Results for %s\n", date)
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=" + pin + "&date=" + date
	return callCowin(url)
}

func callCowinUsingDid(did string, date string) *SessionResponse {
	fmt.Printf("Results for  %s\n", date)
	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=" + did + "&date=" + date
	return callCowin(url)
}

func callCowin(url string) *SessionResponse {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("accept", "application/json, text/plain, */*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
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
				if !contains(BLACKLIST, covidData.Centers[i].CenterID) &&
					covidData.Centers[i].Sessions[j].AvailableCapacity > 0 { // remove unusable centers
					returnData := SessionResponse{covidData.Centers[i], covidData.Centers[i].Sessions[j]}
					return &returnData
				}
			}
		}
	}
	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
