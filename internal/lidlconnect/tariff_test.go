package lidlconnect_test

import (
	"testing"

	"github.com/avakarev/go-testutil"
	"github.com/jarcoal/httpmock"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func TestGetBookedTariff(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.test.host/api/graphql",
		httpmock.NewStringResponder(200, string(testutil.FixtureBytes(t, "./test/fixtures/booked_tariff.json"))))

	tariff, err := lidlconnect.TestClient().GetBookedTariff()
	if err != nil {
		t.Errorf("Failed to get tariff: %s", err.Error())
	}

	testutil.Diff(lidlconnect.BookedTariff{
		TariffID:   "128",
		Name:       "Data S",
		BasicFee:   299,
		Cancelable: true,
	}, *tariff, t)
}
