package rbh

import (
	"rbh/forms"
	"time"
)

type Account struct {
	usn string // username       e.g. zaydek
	acc string // account number e.g. .../5QR43284/
	tok string // token          e.g. 6b1454041e27fa2428ddff8a3467f8138445e7f3
}

func (a *Account) Username() string { return a.usn }
func (a *Account) Token() string    { return a.tok }

// user, err := rbh.Login(username, password)
func Login(username, password string) (*Account, error) {
	// i
	tok, err := getToken(5*time.Second, username, password)
	if err != nil {
		return nil, err
	}
	// ii
	acc, err := getAccNo(5*time.Second, tok)
	if err != nil {
		return nil, err
	}
	// iii
	return &Account{username, acc, tok}, nil
}

func getToken(timeout time.Duration, username, password string) (string, error) {
	// i
	form := forms.Token{username, password}
	// i
	var token types.Token
	if err := http2.Post(timeout, ep.Token, form.Encode(), &token); err != nil {
		return "", err
	}
	// ii
	return token.Token, nil
}

func getAccNo(d time.Duration, token string) (string, error) {
	// i
	var accounts types.Accounts
	if err := http2.AuthGet(d, ep.Accounts, token, &accounts); err != nil {
		return "", err
	}
	// ii
	return accounts.Results[0].URL, nil
}
