package fpl

func (c *Client) ListTransfers(managerID string) ([]TransferHistory, error) {

	url := "https://draft.premierleague.com/api/draft/league/" + managerID + "/transactions"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	var transfer []TransferHistory

	_, err = c.Do(response, &transfer)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
