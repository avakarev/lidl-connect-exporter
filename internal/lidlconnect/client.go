package lidlconnect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

// DefaultAPIHost is lidl-connect api host
const DefaultAPIHost = "api.lidl-connect.de"

var (
	username string
	password string
	host     string
)

// Client implements api client
type Client struct {
	Username    string
	Password    string
	Host        string
	AccessToken string
	mutex       sync.Mutex
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("https://%s/api/%s", c.Host, path)
}

func (c *Client) post(path string, payload interface{}, headers map[string]string) (*http.Response, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(bodyBytes)
	req, err := http.NewRequest(http.MethodPost, c.url(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return http.DefaultClient.Do(req)
}

func (c *Client) auth() map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.AccessToken),
	}
}

func (c *Client) login() error {
	t, err := c.GetToken()
	if err != nil {
		log.Error().Err(err).Msg("login failed")
		return err
	}
	c.AccessToken = t.AccessToken
	log.Debug().Msg("login succeded")
	return nil
}

func (c *Client) graphql(query interface{}) (*http.Response, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	resp, err := c.post("graphql", query, c.auth())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		if err := c.login(); err != nil {
			return nil, err
		}
	}
	return c.post("graphql", query, c.auth())
}

// NewClient returns new client value with given credentials
func NewClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
		Host:     host,
	}
}

// DefaultClient returns new client value with credentials from env
func DefaultClient() *Client {
	return NewClient(username, password)
}

// TestClient returns new mocked client for testing purposes
func TestClient() *Client {
	return &Client{
		Username: "foo",
		Password: "bar",
		Host:     "api.test.host",
	}
}

func init() {
	username = os.Getenv("LIDL_CONNECT_USERNAME")
	password = os.Getenv("LIDL_CONNECT_PASSWORD")
	host = DefaultAPIHost
	if h := os.Getenv("LIDL_CONNECT_HOST"); h != "" {
		host = h
	}
}
