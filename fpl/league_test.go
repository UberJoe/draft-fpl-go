package fpl

import (
	"testing"
)

func TestLeagueDetails(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.GetLeagueDetails("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}

func TestGetStandings(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.GetStandings("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}

// func TestGetNewEntries(t *testing.T) {

// 	c := NewClient(nil)

// 	got, _ := c.GetNewEntries("")
// 	if got != nil {
// 		t.Errorf("Got: %+v", got)
// 	}

// }

// func TestGetTeamInLeague(t *testing.T) {

// 	c := NewClient(nil)

// 	got, _ := c.GetTeamInfoInLeague("", "")
// 	if got != nil {
// 		t.Errorf("Got: %+v", got)
// 	}

// }
