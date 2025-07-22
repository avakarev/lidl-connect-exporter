// Package lidlconnect implements lidl-connect apis
package lidlconnect

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
	q := map[string]any{
		"operationName": "balanceInfo",
		"query": `query
			balanceInfo {
				currentCustomer {
					balance
				}
			}`,
		"variables": "{}",
	}
	var resp BalanceInfoResponse
	return &resp.Data, c.graphql(q, &resp)
}
