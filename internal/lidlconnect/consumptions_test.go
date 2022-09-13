package lidlconnect_test

import (
	"testing"

	"github.com/avakarev/go-testutil"
	"github.com/jarcoal/httpmock"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func TestGetConsumptions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.test.host/api/graphql",
		httpmock.NewStringResponder(200, string(testutil.FixtureBytes(t, "./test/fixtures/consumptions.json"))))

	consumptions, err := lidlconnect.TestClient().GetConsumptions()
	if err != nil {
		t.Errorf("Failed to get consumptions: %s", err.Error())
	}

	testutil.Diff([]lidlconnect.ConsumptionsForUnit{{
		Consumed:       1.05,
		Unit:           "GB",
		FormattedUnit:  "GB",
		Type:           "DATA",
		Description:    "1.05 GB von insgesamt 7.73 GB verbraucht",
		ExpirationDate: "2022-09-29T12:51:12+02:00",
		Left:           6.68,
		Max:            7.73,
		TariffOrOptions: []lidlconnect.TariffOrOption{{
			ID:   "CCS_92037",
			Name: "7 GB Geschenk",
			Type: "Tariffoption",
			Consumptions: []lidlconnect.ConsumptionsForUnit{{
				Consumed:       0.32,
				Unit:           "GB",
				FormattedUnit:  "GB",
				Type:           "DATA",
				Description:    "0.32 GB von insgesamt 7 GB verbraucht",
				ExpirationDate: "2022-09-29T12:51:12+02:00",
				Left:           6.68,
				Max:            7,
			}},
		}, {
			ID:   "128",
			Name: "Data S",
			Type: "Tariff",
			Consumptions: []lidlconnect.ConsumptionsForUnit{{
				Consumed:       750,
				Unit:           "MB",
				FormattedUnit:  "MB",
				Type:           "DATA",
				Description:    "750 MB von insgesamt 750 MB verbraucht",
				ExpirationDate: "2022-09-20T00:00:00+02:00",
				Max:            750,
			}},
		}},
	}}, consumptions, t)
}
