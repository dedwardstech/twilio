package sms

import "net/http"

type MessageParser interface {
	FromRequest(r *http.Request) InboundMessage
}

func NewMessageParser() MessageParser {
	return &parser{}
}

type parser struct {}

func (p parser) FromRequest(r * http.Request) InboundMessage {
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
