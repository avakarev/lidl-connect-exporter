// Package lidlconnect implements lidl-connect apis
package lidlconnect

import (
	"encoding/json"
	"fmt"
	"io"
)

// Customer represents customer attributes
type Customer struct {
	Balance int64 `json:"balance"`
}

// BalanceInfo represents customer's status
type BalanceInfo struct {
	CurrentCustomer Customer `json:"currentCustomer"`
}

// BalanceInfoResponse represents response of `balanceInfo` query
type BalanceInfoResponse struct {
	Data BalanceInfo `json:"data"`
}

// GetBalanceInfo returns current state of the customer's balance
func (c *Client) GetBalanceInfo() (*BalanceInfo, error) {
	q := map[string]string{
		"operationName": "balanceInfo",
		"query": `query
			balanceInfo {
				currentCustomer {
					balance
				}
			}`,
		"variables": "{}",
	}
	res, err := c.graphql(q)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got unexpected status code %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response BalanceInfoResponse
	return &response.Data, json.Unmarshal(bodyBytes, &response)
}
