package lidlconnect

import (
	"fmt"
	"net/http"
	"os"

	"github.com/avakarev/go-util/httputil"
)

// DefaultAPIHost is lidl-connect api host
const DefaultAPIHost = "api.lidl-connect.de"

// Client implements api client
type Client struct {
	base        httputil.Client
	Account     *Account
	AccessToken string
}

func (c *Client) login() error {
	t, err := c.GetToken()
	if err != nil {
		return err
	}
	c.base.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.AccessToken))
	return nil
}

func (c *Client) graphql(payload any, destPtr any) error {
	status, err := c.base.PostJSON("/api/graphql", payload, destPtr)
	if err != nil {
		return err
	}
	if status == http.StatusUnauthorized {
		if err := c.login(); err != nil {
			return err
		}
		// re-try graphql request
		status, err = c.base.PostJSON("/api/graphql", payload, destPtr)
		if err != nil {
			return err
		}
	}
	return httputil.ExpectStatus(200, status)
}

// NewClient returns new client value with given credentials
// Host is optional arg, when not provided, DefaultAPIHost is used
func NewClient(acc *Account, opts ...string) *Client {
	host := DefaultAPIHost
	if len(opts) > 0 && opts[0] != "" {
		host = opts[0]
	}
	if host == DefaultAPIHost {
		if h := os.Getenv("LIDL_CONNECT_HOST"); h != "" {
			host = h
		}
	}
	return &Client{
		base: httputil.Client{
			BaseURL: "https://" + host,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		Account: acc,
	}
}
