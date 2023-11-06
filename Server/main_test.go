package main

import (
	"strconv"
	"testing"
)

func TestGetEmails(t *testing.T) {
	// Generate test array
	var commits []gitHubCommit
	for i := 0; i < 3; i++ {
		var singleCommit gitHubCommit
		var generatedEmail string = strconv.Itoa(i) + "@gmail.com"
		singleCommit.Author = gitHubAuthor{Email: generatedEmail}
		commits = append(commits, singleCommit)
	}
	commits = append(commits, commits[0])
	if len(commits) != 4 {
		t.Errorf("Array should contains 4 objects")
	}

	// Call getEmails function
	var emails = getEmails(commits)
	if len(emails) != 3 {
		t.Errorf("Array should contains 4 objects")
	}
}

func TestContains(t *testing.T) {
	var testArray []string = []string{"1", "2", "3"}
	var testCaseTrue string = "1"
	var testCaseFalse string = "4"

	if !contains(testArray, testCaseTrue) {
		t.Errorf("Array %s should contains %s", testArray, testCaseTrue)
	}

	if contains(testArray, testCaseFalse) {
		t.Errorf("Array %s should not contains %s", testArray, testCaseFalse)
	}
}

func TestParseEvents(t *testing.T) {
	// Generate test data
	var events []gitHubEvent
	for i := 0; i < 3; i++ {
		singleEvent := gitHubEvent{
			Id:   strconv.Itoa(i),
			Type: "type" + strconv.Itoa(i),
			Actor: gitHubActor{
				Login: "login" + strconv.Itoa(i),
			},
			Repo: gitHubRepo{
				Name: "name" + strconv.Itoa(i),
			},
			Created_at: "some time...",
		}
		events = append(events, singleEvent)
	}

	// Call parseEvents function
	var resEvents = parseEvents((events))
	if len(resEvents.Events) != 3 {
		t.Errorf("Events list contains 3 objects")
	}
	if len(resEvents.Actors) != 3 {
		t.Errorf("Actor list contains 3 objects")
	}
	if len(resEvents.Repos) != 3 {
		t.Errorf("Repos list contains 3 objects")
	}
	if len(resEvents.Emails) != 0 {
		t.Errorf("Emails list contains 0 objects")
	}
}
