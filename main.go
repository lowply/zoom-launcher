package main

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Email    string
	Regex    string
	CredsDir string
}

var config Config

func main() {
	if os.Getenv("ZL_EMAIL") == "" {
		log.Fatal("ZL_EMAIL is empty")
	}

	if os.Getenv("ZL_REGEX") == "" {
		log.Fatal("ZL_REGEX is empty")
	}

	config.Email = os.Getenv("ZL_EMAIL")
	config.Regex = os.Getenv("ZL_REGEX")
	config.CredsDir = os.Getenv("HOME") + "/.config/google-calendar-api"

	_, err := os.Stat(config.CredsDir + "/credentials.json")
	if err != nil {
		log.Fatal(err)
	}

	g := NewGcal()

	events, err := g.listEvents()
	if err != nil {
		log.Fatal(err)
	}

	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
		os.Exit(0)
	}

	meetings := []*meeting{}

	// Create a meeting only when I accepted it
	for _, item := range events.Items {
		for _, attendee := range item.Attendees {
			if attendee.Email == config.Email && attendee.ResponseStatus == "accepted" {
				m, err := NewMeeting(item)
				if err != nil {
					log.Fatal(err)
				}
				meetings = append(meetings, m)
			}
		}
	}

	// Show next 3 meetings
	for i, m := range meetings {
		u := ""
		if m.url == nil {
			u = "No meeting URL"
		} else {
			u = m.url.String()
		}
		fmt.Printf("[%v] %v: %v\n", m.date, m.name, u)

		if i == 2 {
			break
		}
	}

	// Open the next meeting URL
	err = meetings[0].open()
	if err != nil {
		log.Println(err)
	}
}
