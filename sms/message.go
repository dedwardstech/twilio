package sms

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
	return s.v["From"]
}

func (s *sender) Country() string {
	return s.v["FromCountry"]
}

func (s *sender) State() string {
	return s.v["FromState"]
}

func (s *sender) City() string {
	return s.v["FromCity"]
}

func (s *sender) Zip() string {
	return s.v["FromZip"]
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
	return s.v["To"]
}

func (s *receiver) Country() string {
	return s.v["ToCountry"]
}

func (s *receiver) State() string {
	return s.v["ToState"]
}

func (s *receiver) City() string {
	return s.v["ToCity"]
}

func (s *receiver) Zip() string {
	return s.v["ToZip"]
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
	return m.v["SmsMessageSid"]
}

func (m *message) Sid() string {
	return m.v["SmsSid"]
}

func (m *message) AccountSid() string {
	return m.v["AccountSid"]
}

func (m *message) MediaCount() string {
	return m.v["NumMedia"]
}

func (m *message) SegmentCount() string {
	return m.v["NumSegments"]
}

func (m *message) Body() string {
	return m.v["Body"]
}

func (m *message) Status() string {
	return m.v["SmsStatus"]
}

func (m *message) APIVersion() string {
	return m.v["APIVersion"]
}