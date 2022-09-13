package lidlconnect

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog/log"
)

// TokenClaim represents client credentials
type TokenClaim struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

// NewTokenClaim returns client credentials for login
func NewTokenClaim(username string, password string) *TokenClaim {
	return &TokenClaim{
		ClientID:     "lidl",
		ClientSecret: "lidl",
		GrantType:    "password",
		Username:     username,
		Password:     password,
	}
}

// Token represents api auth token
type Token struct {
	Type         string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RequestedAt  time.Time
}

// Expired checks whether token is expired
func (t *Token) Expired() bool {
	return time.Since(t.RequestedAt).Seconds() >= float64(t.ExpiresIn)
}

// GetToken return access token for the given credentials
func (c *Client) GetToken() (*Token, error) {
	if c.Username == "" || c.Password == "" {
		log.Fatal().Msg("LIDL_CONNECT_USERNAME and/or LIDL_CONNECT_PASSWORD is empty")
	}
	claim := NewTokenClaim(c.Username, c.Password)
	res, err := c.post("token", claim, nil)
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

	token := Token{RequestedAt: time.Now()}
	return &token, json.Unmarshal(bodyBytes, &token)
}
