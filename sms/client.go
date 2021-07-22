package sms

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type ClientOptions struct {
	AccountSid string
	AuthToken  string
	Number     string
}

type Client interface {
	NewRequest(to, msg string) (*http.Request, error)
}

func NewClient(options *ClientOptions) (Client, error) {
	if options.AccountSid == "" {
		return nil, errors.New("accountSid required")
	}

	if options.AuthToken == "" {
		return nil, errors.New("authToken required")
	}

	if options.Number == "" {
		return nil, errors.New("number required")
	}

	if !isValidPhoneNumber(options.Number) {
		return nil, errors.New("valid phone number required")
	}

	return &client{
		opts: options,
		url:  fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", options.AccountSid),
	}, nil
}

type client struct {
	opts *ClientOptions
	url  string
}

func (c client) NewRequest(to, msg string) (*http.Request, error) {
	msgDataReader := strings.NewReader(c.buildMsgData(to, msg))
	req, err := http.NewRequest("POST", c.url, msgDataReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.opts.AccountSid, c.opts.AuthToken)

	return req, nil
}

func (c client) buildMsgData(to, msg string) string {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", c.opts.Number)
	msgData.Set("Body", msg)

	return msgData.Encode()
}

var (
	phoneRegex = regexp.MustCompile(`^\+[0-9]{11}$`)
)

func isValidPhoneNumber(number string) bool {
	return phoneRegex.Match([]byte(number))
}
