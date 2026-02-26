package lidlconnect

import (
	"fmt"
	"strings"
)

// Account holds account credentials
type Account struct {
	Name     string
	Username string
	Password string // #nosec G117
}

// Validate checks account validity
func (acc *Account) Validate() error {
	if acc.Username == "" {
		return fmt.Errorf("username is required but missing")
	}
	if acc.Password == "" {
		return fmt.Errorf("password is required but missing")
	}
	if acc.Name == "" {
		return fmt.Errorf("name is required but missing")
	}
	return nil
}

// ParseAccount parses account string into struct
func ParseAccount(accStr string) (*Account, error) {
	name, creds, ok := strings.Cut(accStr, "=")
	if !ok {
		return nil, fmt.Errorf("bad format: expected name=usr:pwd")
	}

	usr, pwd, ok := strings.Cut(creds, ":")
	if !ok {
		return nil, fmt.Errorf("bad format: expected name=usr:pwd")
	}

	acc := &Account{
		Name:     name,
		Username: usr,
		Password: pwd,
	}

	if err := acc.Validate(); err != nil {
		return nil, err
	}
	return acc, nil
}
