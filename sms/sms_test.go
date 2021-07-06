package sms

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	is2 "github.com/matryer/is"
)

func TestParseFromHttpRequest(t *testing.T) {
	is := is2.New(t)

	v := url.Values{}
	v.Add("From", "4048675309")
	v.Add("FromCountry", "USA")
	v.Add("FromState", "GA")
	v.Add("FromCity", "Atlanta")
	v.Add("FromZip", "30000")

	v.Add("To", "7708675309")
	v.Add("ToCountry", "US")
	v.Add("ToState", "FL")
	v.Add("ToCity", "Miami")
	v.Add("ToZip", "40000")

	v.Add("SmsMessageSid", "smsmessagesid")
	v.Add("SmsSid", "smssid")
	v.Add("AccountSid", "accountsid")
	v.Add("NumMedia", "1")
	v.Add("NumSegments", "2")
	v.Add("Body", "body!")
	v.Add("SmsStatus", "received")
	v.Add("ApiVersion", "2010-04-01")

	r := httptest.NewRequest("POST", "/", strings.NewReader(v.Encode()))

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	err := r.ParseForm()
	if err != nil {
		t.Errorf("test setup failed: %v", err)
		return
	}

	m := ParseFromHttpRequest(r)


	is.Equal(m.Receiver().PhoneNumber(), "7708675309")
	is.Equal(m.Receiver().Country(), "US")
	is.Equal(m.Receiver().State(), "FL")
	is.Equal(m.Receiver().City(), "Miami")
	is.Equal(m.Receiver().Zip(), "40000")

	is.Equal(m.Sender().PhoneNumber(), "4048675309")
	is.Equal(m.Sender().Country(), "USA")
	is.Equal(m.Sender().State(), "GA")
	is.Equal(m.Sender().Zip(), "30000")

	is.Equal(m.MessageSid(), "smsmessagesid")
	is.Equal(m.Sid(), "smssid")
	is.Equal(m.AccountSid(), "accountsid")
	is.Equal(m.MediaCount(), "1")
	is.Equal(m.SegmentCount(), "2")
	is.Equal(m.Body(), "body!")
	is.Equal(m.Status(), "received")
	is.Equal(m.ApiVersion(), "2010-04-01")
}