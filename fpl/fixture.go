package fpl

import (
	"strconv"
)

const ()

func (c *Client) GetWeeklyFixture(gameWeek int) (*WeeklyFixture, error) {
	url := "https://draft.premierleague.com/api/event/" + strconv.Itoa(gameWeek) + "/fixtures"

	wf := &WeeklyFixture{}

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(response, wf)
	if err != nil {
		return nil, err
	}

	return wf, nil
}
