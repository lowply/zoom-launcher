package main

import (
	"fmt"
	"log"
	"os"
)

const (
	CredsDir = "/.config/google-calendar-api"
	ConfPath = "/.config/zoom-launcher.yaml"
)

func main() {
	_, err := os.Stat(os.Getenv("HOME") + CredsDir + "/credentials.json")
	if err != nil {
		log.Fatal(err)
	}

	err = readconfig(os.Getenv("HOME") + ConfPath)
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
	for _, m := range meetings[:3] {
		u := ""
		if m.url == nil {
			u = "No meeting URL"
		} else {
			u = m.url.String()
		}
		fmt.Printf("[%v] %v: %v\n", m.date, m.name, u)
	}

	// Open the next meeting URL
	err = meetings[0].open()
	if err != nil {
		log.Println(err)
	}
}
