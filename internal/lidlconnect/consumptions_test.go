package lidlconnect_test

import (
	"testing"
	"time"

	"github.com/avakarev/go-util/testutil"
	"github.com/avakarev/go-util/timeutil"
	"github.com/jarcoal/httpmock"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func TestGetConsumptions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.test.host/api/graphql",
		httpmock.NewStringResponder(200, string(testutil.FixtureBytes(t, "./test/fixtures/consumptions.json"))))

	acc := &lidlconnect.Account{Username: "usr", Password: "pwd", Name: "test"}
	consumptions, err := lidlconnect.NewClient(acc, "api.test.host").GetConsumptions()
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

func TestConsumptionsForUnitExpiresInSecondsFrom(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.test.host/api/graphql",
		httpmock.NewStringResponder(200, string(testutil.FixtureBytes(t, "./test/fixtures/consumptions.json"))))

	acc := &lidlconnect.Account{Username: "usr", Password: "pwd", Name: "test"}
	consumptions, err := lidlconnect.NewClient(acc, "api.test.host").GetConsumptions()
	if err != nil {
		t.Errorf("Failed to get consumptions: %s", err.Error())
	}

	testutil.Diff(1, len(consumptions), t)

	correctTime, _ := time.ParseInLocation(time.RFC3339, "2022-09-15T01:22:21+02:00", timeutil.Location)
	testutil.Diff(1250931.0, consumptions[0].ExpiresIn(correctTime).Seconds(), t)

	wrongTime, _ := time.ParseInLocation(time.RFC3339, "2022-09-29T12:51:13+02:00", timeutil.Location)
	testutil.Diff(0.0, consumptions[0].ExpiresIn(wrongTime).Seconds(), t)
}
