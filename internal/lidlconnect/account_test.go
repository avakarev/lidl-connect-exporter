package lidlconnect_test

import (
	"errors"
	"testing"

	"github.com/avakarev/go-util/testutil"

	"github.com/avakarev/lidl-connect-exporter/internal/lidlconnect"
)

func TestParseAccount(t *testing.T) {
	cases := []struct {
		str string
		acc *lidlconnect.Account
		err error
	}{
		{str: "", acc: nil, err: errors.New("bad format: expected name=usr:pwd")},
		{str: "foobar", acc: nil, err: errors.New("bad format: expected name=usr:pwd")},
		{str: "usr:pwd", acc: nil, err: errors.New("bad format: expected name=usr:pwd")},
		{str: "alias=", acc: nil, err: errors.New("bad format: expected name=usr:pwd")},
		{str: "alias=usr:pwd", acc: &lidlconnect.Account{Name: "alias", Username: "usr", Password: "pwd"}, err: nil},
	}
	for _, tcase := range cases {
		acc, err := lidlconnect.ParseAccount(tcase.str)
		if tcase.err != nil {
			testutil.MustErr(tcase.err, err, t)
		} else {
			testutil.MustNoErr(err, t)
		}
		if tcase.acc != nil {
			testutil.Diff(*tcase.acc, *acc, t)
		} else {
			testutil.Diff(true, acc == nil, t)
		}
	}
}
