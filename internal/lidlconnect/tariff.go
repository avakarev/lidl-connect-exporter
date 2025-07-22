package lidlconnect

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
	q := map[string]any{
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
	var resp TariffsResponse
	return &resp.Data.Tariffs.BookedTariff, c.graphql(q, &resp)
}
