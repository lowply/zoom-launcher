package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	calendar "google.golang.org/api/calendar/v3"
)

type meeting struct {
	name    string
	date    time.Time
	url     *url.URL
	zoomUrl *url.URL
}

func NewMeeting(item *calendar.Event) (*meeting, error) {
	m := &meeting{
		name: item.Summary,
	}

	date := item.Start.DateTime
	if date == "" {
		date = item.Start.Date
	}

	d, err := time.Parse("2006-01-02T15:04:05-07:00", date)
	if err != nil {
		return nil, err
	}
	m.date = d

	c := item.ConferenceData

	if c == nil {
		// No conference has been set.
		// Checking the description field to see if there's the Zoom URL
		lines := strings.Split(item.Description, "<br>")
		for _, d := range lines {
			re := regexp.MustCompile(config.Regex)
			find := re.Find([]byte(d))
			if len(find) != 0 {
				m.url, err = url.Parse(string(find))
				if err != nil {
					return nil, err
				}
			}
		}
	} else if c.ConferenceSolution.Name != "Zoom Meeting" {
		// Conference has been set, but not a Zoom meeting
	} else {
		// A Zoom conference has been set.
		for _, e := range c.EntryPoints {
			matched, err := regexp.Match(config.Regex, []byte(e.Uri))
			if err != nil {
				return nil, err
			}
			if matched {
				m.url, err = url.Parse(e.Uri)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if m.url != nil {
		ru := []byte(m.url.RequestURI())
		re := regexp.MustCompile(`\?`)
		ru = re.ReplaceAll(ru, []byte("&"))
		re = regexp.MustCompile(`/j/`)
		ru = re.ReplaceAll(ru, []byte("/join?confno="))
		m.zoomUrl, err = url.Parse("zoommtg://" + m.url.Hostname() + string(ru))
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *meeting) open() error {
	if m.zoomUrl == nil {
		return nil
	}
	fmt.Printf("Do you want to join \"%v\" now? y/n\n", m.name)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		cmd := exec.Command("open", m.zoomUrl.String())
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
