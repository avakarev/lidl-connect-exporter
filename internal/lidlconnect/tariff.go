package lidlconnect

import (
	"encoding/json"
	"fmt"
	"io"
)

// BookedTariff represents tariff attributes
type BookedTariff struct {
	TariffID   string `json:"tariffId"`
	Name       string `json:"name"`
	BasicFee   int64  `json:"basicFee"`
	Cancelable bool   `json:"cancelable"`
}

// Tariffs represents set of assigned tariffs
type Tariffs struct {
	BookedTariff BookedTariff `json:"bookedTariff"`
}

// TariffsData is Tariffs envelope
type TariffsData struct {
	Tariffs Tariffs `json:"tariffs"`
}

// TariffsResponse represents response of `tariffs` query
type TariffsResponse struct {
	Data TariffsData `json:"data"`
}

// GetBookedTariff returns booked tariff
func (c *Client) GetBookedTariff() (*BookedTariff, error) {
	q := map[string]interface{}{
		"operationName": nil,
		"query": `query
			bookedTariff {
				tariffs {
					bookedTariff {
						tariffId
						name
						basicFee
						cancelable
					}
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

	var response TariffsResponse
	return &response.Data.Tariffs.BookedTariff, json.Unmarshal(bodyBytes, &response)
}
