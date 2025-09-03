package fpl

import (
	"fmt"
	"strconv"
)

func (c *Client) GetWeeklyFixture(gameWeek int) (*WeeklyFixture, error) {
	if gameWeek <= 0 {
		return nil, fmt.Errorf("invalid game week: %d", gameWeek)
	}

	url := "https://draft.premierleague.com/api/event/" + strconv.Itoa(gameWeek) + "/fixtures"

	wf := &WeeklyFixture{}

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for weekly fixture: %w", err)
	}

	_, err = c.Do(response, wf)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly fixture: %w", err)
	}

	return wf, nil
}
