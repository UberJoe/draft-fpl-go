package fpl

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// points of the manager by given game week
func (c *Client) ListWeeklyPoints(gameWeek int, managerID string) ([]EntryHistoryWeek, error) {
	if gameWeek <= 0 {
		return nil, fmt.Errorf("invalid game week: %d", gameWeek)
	}

	url := "https://draft.premierleague.com/api/entry/" + managerID + "/history"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for manager history: %w", err)
	}

	v := &EntryHistory{}

	_, err = c.Do(response, v)
	if err != nil {
		return nil, err
	}

	var wResponse []EntryHistoryWeek
	for index, value := range v.History {

		if value.Event == gameWeek {

			m, err := json.Marshal(v.History[index])
			if err != nil {
				return nil, err
			}

			w := &EntryHistoryWeek{}

			if err := json.Unmarshal(m, &w); err != nil {
				return nil, err
			}

			wResponse = append(wResponse, *w)
			return wResponse, nil
		}
	}
	return nil, errors.New("could not find game week")
}

// points of the manager for all game weeks
func (c *Client) ListAllWeeks(managerID string) ([]EntryHistoryWeek, error) {

	url := "https://draft.premierleague.com/api/entry/" + managerID + "/history"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	v := &EntryHistory{}

	_, err = c.Do(response, v)
	if err != nil {
		return nil, err
	}

	return v.History, nil
}

// individual performances of players by given game week
func (c *Client) ListWeeklyPerformance(gameWeek int, managerID string) ([]TeamWeekly, error) {

	url := "https://draft.premierleague.com/api/entry/" + managerID + "/event/" + strconv.Itoa(gameWeek)

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	t := &TeamWeekly{}

	_, err = c.Do(response, t)
	if err != nil {
		return nil, err
	}

	var teamweekly []TeamWeekly
	teamweekly = append(teamweekly, *t)

	return teamweekly, nil
}
