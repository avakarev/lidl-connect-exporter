package lidlconnect

import (
	"time"

	"github.com/avakarev/go-util/timeutil"
)

// TariffOrOption represents tariff option attributes
type TariffOrOption struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Type         string                `json:"type"`
	Consumptions []ConsumptionsForUnit `json:"consumptions"`
}

// ConsumptionsForUnit represents consumptions attributes
type ConsumptionsForUnit struct {
	Consumed        float64          `json:"consumed"`
	Unit            string           `json:"unit"`
	FormattedUnit   string           `json:"formattedUnit"`
	Type            string           `json:"type"`
	Description     string           `json:"description"`
	ExpirationDate  string           `json:"expirationDate"`
	Left            float64          `json:"left"`
	Max             float64          `json:"max"`
	TariffOrOptions []TariffOrOption `json:"tariffOrOptions"`
}

// ExpiresIn return consumption expiration duration
func (c ConsumptionsForUnit) ExpiresIn(t time.Time) time.Duration {
	expiresAt, err := time.ParseInLocation(time.RFC3339, c.ExpirationDate, timeutil.Location)
	if err != nil || expiresAt.Before(t) {
		return time.Duration(0)
	}
	return expiresAt.Sub(t)
}

// Consumptions represents current consumptions
type Consumptions struct {
	ConsumptionsForUnit []ConsumptionsForUnit `json:"consumptionsForUnit"`
}

// ConsumptionsData is Consumptions envelope
type ConsumptionsData struct {
	Consumptions Consumptions `json:"consumptions"`
}

// ConsumptionsResponse represents response of `consumptions` query
type ConsumptionsResponse struct {
	Data ConsumptionsData `json:"data"`
}

// GetConsumptions returns current consumptions
func (c *Client) GetConsumptions() ([]ConsumptionsForUnit, error) {
	q := map[string]any{
		"operationName": "consumptions",
		"query": `query
			consumptions {
				consumptions {
					consumptionsForUnit {
						consumed
						unit
						formattedUnit
						type
						description
						expirationDate
						left
						max
						tariffOrOptions {
							name
							id
							type
							consumptions {
								consumed
								unit
								formattedUnit
								type
								description
								expirationDate
								left
								max
							}
						}
					}
				}
			}`,
		"variables": "{}",
	}
	var resp ConsumptionsResponse
	return resp.Data.Consumptions.ConsumptionsForUnit, c.graphql(q, &resp)
}
