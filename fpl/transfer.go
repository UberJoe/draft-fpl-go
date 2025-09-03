package fpl

import "fmt"

func (c *Client) ListTransfers(managerID string) ([]TransferHistory, error) {
	if managerID == "" {
		return nil, fmt.Errorf("managerID cannot be empty")
	}

	url := "https://draft.premierleague.com/api/draft/league/" + managerID + "/transactions"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for transfer history: %w", err)
	}

	var transfer []TransferHistory

	_, err = c.Do(response, &transfer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transfer history: %w", err)
	}

	return transfer, nil
}
