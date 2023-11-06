package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gitHubEvent struct {
	Id         string
	Type       string
	Actor      gitHubActor
	Repo       gitHubRepo
	Payload    gitHubPayload
	Created_at string
}

type gitHubActor struct {
	Login string
}

type gitHubRepo struct {
	Name string
}

type gitHubPayload struct {
	Commits []gitHubCommit
}

type gitHubCommit struct {
	Author gitHubAuthor
}

type gitHubAuthor struct {
	Email string
}

type event struct {
	Id      string
	Type    string
	Actor   string
	RepoUrl string
	Emails  []string
	Created string
}

type responseJson struct {
	Events []event
	Actors []string //Up to 50
	Repos  []string //Up to 20
	Emails []string
}

// should store the last 50 actoers
// should store the last 20 repositories
// should store all mails

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	gitHubEvents, err := getEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	events := parseEvents(gitHubEvents)
	resJson, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resJson))
}

func getEvents() ([]gitHubEvent, error) {
	resp, err := http.Get("https://api.github.com/users/facebook/events")
	if err != nil {
		log.Fatal([]byte("Error!\n"))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal([]byte("Error!\n"))
	}

	var events []gitHubEvent
	err = json.Unmarshal([]byte(body), &events)
	if err != nil {
		log.Fatal([]byte("Error!\n"))
	}
	return events, nil
}

func getEmails(commits []gitHubCommit) []string {
	var emails []string
	for i := 0; i < len(commits); i++ {
		singleEmail := commits[i].Author.Email
		if !contains(emails, singleEmail) {
			emails = append(emails, singleEmail)
		}
	}
	return emails
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func parseEvents(rawGitHubEvents []gitHubEvent) responseJson {
	var events []event
	var actors []string
	var repos []string
	var emails []string
	for i := 0; i < len(rawGitHubEvents); i++ {
		singleEvent := event{
			Id:      rawGitHubEvents[i].Id,
			Type:    rawGitHubEvents[i].Type,
			Actor:   rawGitHubEvents[i].Actor.Login,
			RepoUrl: rawGitHubEvents[i].Repo.Name,
			Created: rawGitHubEvents[i].Created_at,
		}
		singleEvent.Emails = getEmails(rawGitHubEvents[i].Payload.Commits)
		events = append(events, singleEvent)

		// Handle Actor
		currentActor := rawGitHubEvents[i].Actor.Login
		if len(actors) < 50 && !contains(actors, currentActor) {
			actors = append(actors, currentActor)
		}
		// Handle Repo
		currentRepo := rawGitHubEvents[i].Repo.Name
		if len(actors) < 20 && !contains(repos, currentRepo) {
			repos = append(repos, currentRepo)
		}
		// Handle Email
		for i := 0; i < len(singleEvent.Emails); i++ {
			if !contains(emails, singleEvent.Emails[i]) {
				emails = append(emails, singleEvent.Emails[i])
			}
		}
	}

	res := responseJson{
		Events: events,
		Actors: actors,
		Repos:  repos,
		Emails: emails,
	}
	return res
}
