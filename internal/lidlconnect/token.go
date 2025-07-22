package lidlconnect

import (
	"time"

	"github.com/avakarev/go-util/httputil"
)

// TokenClaim represents client credentials
type TokenClaim struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

// Token represents api auth token response
type Token struct {
	Type         string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	RequestedAt  time.Time `json:"-"`
}

// Expired checks whether token is expired
func (t *Token) Expired() bool {
	return time.Since(t.RequestedAt).Seconds() >= float64(t.ExpiresIn)
}

// GetToken return access token for the given credentials
func (c *Client) GetToken() (*Token, error) {
	if err := c.Account.Validate(); err != nil {
		return nil, err
	}
	claim := &TokenClaim{
		ClientID:     "lidl",
		ClientSecret: "lidl",
		GrantType:    "password",
		Username:     c.Account.Username,
		Password:     c.Account.Password,
	}

	var resp Token
	status, err := c.base.PostJSON("/api/token", claim, &resp)
	if err != nil {
		return nil, err
	}
	if err := httputil.ExpectStatus(200, status); err != nil {
		return nil, err
	}

	resp.RequestedAt = time.Now()
	return &resp, nil
}
