package lidlconnect_test

import (
	"testing"

	"github.com/avakarev/go-util/testutil"
	"github.com/jarcoal/httpmock"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func TestGetBalanceInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.test.host/api/graphql",
		httpmock.NewStringResponder(200, string(testutil.FixtureBytes(t, "./test/fixtures/balance_info.json"))))

	acc := &lidlconnect.Account{Username: "usr", Password: "pwd", Name: "test"}
	balance, err := lidlconnect.NewClient(acc, "api.test.host").GetBalanceInfo()
	if err != nil {
		t.Errorf("Failed to get balance: %s", err.Error())
	}

	testutil.Diff(lidlconnect.BalanceInfo{
		CurrentCustomer: lidlconnect.Customer{Balance: 701},
	}, *balance, t)
}
