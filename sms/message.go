package sms

const (
	msgSenderPhone   = "From"
	msgSenderCountry = "FromCountry"
	msgSenderState   = "FromState"
	msgSenderCity    = "FromCity"
	msgSenderZip     = "FromZip"

	msgReceiverPhone   = "To"
	msgReceiverCountry = "ToCountry"
	msgReceiverState   = "ToState"
	msgReceiverCity    = "ToCity"
	msgReceiverZip     = "ToZip"

	msgMessageSid   = "SmsMessageSid"
	msgSid          = "SmsSid"
	msgAccountSid   = "AccountSid"
	msgMediaCount   = "NumMedia"
	msgSegmentCount = "NumSegments"
	msgBody         = "Body"
	msgStatus       = "SmsStatus"
	msgAPIVersion   = "ApiVersion"
)

type OutboundMessage struct {
	To   string `json:"to"`
	From string `json:"from"`
	Msg  string `json:"payload"`
}

// Sender provides information, from a Twilio SMS webhook, about the sender of a SMS InboundMessage
type Sender interface {
	PhoneNumber() string
	Country() string
	State() string
	City() string
	Zip() string
}

type sender struct {
	v map[string]string
}

func (s *sender) PhoneNumber() string {
	return s.v[msgSenderPhone]
}

func (s *sender) Country() string {
	return s.v[msgSenderCountry]
}

func (s *sender) State() string {
	return s.v[msgSenderState]
}

func (s *sender) City() string {
	return s.v[msgSenderCity]
}

func (s *sender) Zip() string {
	return s.v[msgSenderZip]
}

// Receiver provides information, from a Twilio SMS webhook, about the receiver of a SMS InboundMessage
type Receiver interface {
	PhoneNumber() string
	Country() string
	State() string
	City() string
	Zip() string
}

type receiver struct {
	v map[string]string
}

func (s *receiver) PhoneNumber() string {
	return s.v[msgReceiverPhone]
}

func (s *receiver) Country() string {
	return s.v[msgReceiverCountry]
}

func (s *receiver) State() string {
	return s.v[msgReceiverState]
}

func (s *receiver) City() string {
	return s.v[msgReceiverCity]
}

func (s *receiver) Zip() string {
	return s.v[msgReceiverZip]
}

// InboundMessage is a SMS message sent from a user to the application
type InboundMessage interface {
	Sender() Sender
	Receiver() Receiver

	MessageSid() string
	Sid() string
	AccountSid() string
	MediaCount() string
	SegmentCount() string
	Body() string
	Status() string

	APIVersion() string
}

type message struct {
	s Sender
	r Receiver
	v map[string]string
}

func (m *message) Sender() Sender {
	return m.s
}

func (m *message) Receiver() Receiver {
	return m.r
}

func (m *message) MessageSid() string {
	return m.v[msgMessageSid]
}

func (m *message) Sid() string {
	return m.v[msgSid]
}

func (m *message) AccountSid() string {
	return m.v[msgAccountSid]
}

func (m *message) MediaCount() string {
	return m.v[msgMediaCount]
}

func (m *message) SegmentCount() string {
	return m.v[msgSegmentCount]
}

func (m *message) Body() string {
	return m.v[msgBody]
}

func (m *message) Status() string {
	return m.v[msgStatus]
}

func (m *message) APIVersion() string {
	return m.v[msgAPIVersion]
}
