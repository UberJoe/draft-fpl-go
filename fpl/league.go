package fpl

import (
	"encoding/json"
	"errors"
)

// league endpoint according to given league id
func (c *Client) GetLeagueDetails(leagueID string) (*LeagueInfo, error) {

	url := "https://draft.premierleague.com/api/league/" + leagueID + "/details"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	l := &LeagueInfo{}

	_, err = c.Do(response, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (c *Client) GetLeague(leagueID string) (*LeagueResponse, error) {
	s, err := c.GetLeagueDetails(leagueID)
	if err != nil {
		return nil, err
	}

	var l *LeagueResponse

	m, err := json.Marshal(s.League)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &l); err != nil {
		return nil, err
	}

	return l, nil

}

func (c *Client) GetLeagueEntries(leagueID string) (*LeagueEntriesResponse, error) {
	s, err := c.GetLeagueDetails(leagueID)
	if err != nil {
		return nil, err
	}

	var e *LeagueEntriesResponse

	m, err := json.Marshal(s.League)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

// standings according to given league id
func (c *Client) GetStandings(leagueID string) ([]StandingsResponse, error) {

	s, err := c.GetLeagueDetails(leagueID)
	if err != nil {
		return nil, err
	}

	standings := s.Standings

	var league []StandingsResponse
	l := &StandingsResponse{}

	for _, v := range standings {
		resp, _ := json.Marshal(v)
		if err = json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}
		league = append(league, *l)

	}

	return league, nil

}

// information about the team according to league
func (c *Client) GetTeamInfoInLeague(leagueID string, entryId int) (*StandingsResponse, error) {

	league, err := c.GetStandings(leagueID)
	if err != nil {
		return nil, err
	}

	for _, v := range league {
		if v.LeagueEntry == entryId {
			return &v, nil
		}
	}

	return nil, errors.New("could not find team in league")
}
