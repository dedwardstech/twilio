package sms

import (
	"fmt"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^\+[0-9]{11}$`)

func IsValidSMSNumber(number string) bool {
	return phoneRegex.Match([]byte(number))
}

func makeSMSUrl(accountSID string) string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSID)
}
