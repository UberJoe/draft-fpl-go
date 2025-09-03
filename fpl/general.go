package fpl

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetGeneral() (*General, error) {

	url := "https://draft.premierleague.com/api/bootstrap-static"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}
	general := &General{}

	_, err = c.Do(response, general)
	if err != nil {
		return nil, err
	}

	return general, nil
}

// information of current EPL clubs
func (c *Client) ListClubsInfo() ([]TeamResponse, error) {

	general, err := c.GetGeneral()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch general data: %w", err)
	}

	var t []TeamResponse

	m, err := json.Marshal(general.Teams)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &t); err != nil {
		return nil, err
	}

	return t, nil
}

// informations of game weeks
func (c *Client) ListEventInfo() (*EventsResponse, error) {

	general, err := c.GetGeneral()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch general data: %w", err)
	}

	var e *EventsResponse

	m, err := json.Marshal(general.Events)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

// element stats
func (c *Client) ListElementStatsInfo() ([]ElementStatsResponse, error) {

	general, _ := c.GetGeneral()

	var e []ElementStatsResponse

	m, err := json.Marshal(general.ElementStats)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

// element types(positions)
func (c *Client) ListElementTypesInfo() ([]ElementTypesResponse, error) {

	general, _ := c.GetGeneral()

	var e []ElementTypesResponse

	m, err := json.Marshal(general.ElementTypes)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

// game settings
func (c *Client) ListSettings() ([]SettingsResponse, error) {

	general, _ := c.GetGeneral()

	g := &SettingsResponse{}

	m, err := json.Marshal(general.Settings)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &g); err != nil {
		return nil, err
	}

	var gameResponse []SettingsResponse
	gameResponse = append(gameResponse, *g)

	return gameResponse, nil

}
