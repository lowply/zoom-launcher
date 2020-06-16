package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

// Mostly taken from https://developers.google.com/calendar/quickstart/go

type gcal struct {
	token      string
	credential string
}

func NewGcal() gcal {
	return gcal{
		token:      config.CredsDir + "/token.json",
		credential: config.CredsDir + "/credentials.json",
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func (g *gcal) getClient(config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := g.tokenFromFile(g.token)
	if err != nil {
		tok, err = g.getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		err := g.saveToken(g.token, tok)
		if err != nil {
			return nil, err
		}
	}
	return config.Client(context.Background(), tok), nil
}

// Request a token from the web, then returns the retrieved token.
func (g *gcal) getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Printf("Unable to read authorization code: %v", err)
		return nil, err
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Printf("Unable to retrieve token from web: %v", err)
		return nil, err
	}
	return tok, nil
}

// Retrieves a token from a local file.
func (g *gcal) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func (g *gcal) saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("Unable to cache oauth token: %v", err)
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func (g *gcal) listEvents() (*calendar.Events, error) {
	b, err := ioutil.ReadFile(g.credential)
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	client, err := g.getClient(config)
	if err != nil {
		return nil, err
	}

	srv, err := calendar.New(client)
	if err != nil {
		log.Printf("Unable to retrieve Calendar client: %v", err)
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Printf("Unable to retrieve next ten of the user's events: %v", err)
		return nil, err
	}

	return events, nil
}
