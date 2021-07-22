package sms

import (
	"net/http"
)

// FromRequest ...
func FromRequest(r *http.Request) InboundMessage {
	msg := make(map[string]string)

	for key, values := range r.PostForm {
		if values == nil || len(values) == 0 {
			msg[key] = ""
		} else {
			msg[key] = values[0]
		}
	}

	return &message{
		s: &sender{v: msg},
		r: &receiver{v: msg},
		v: msg,
	}
}
