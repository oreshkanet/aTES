package mq

import "fmt"

type Validator interface {
	ValidateBytes(subject []byte, domain string, event string, version string) error
	ValidateString(subject string, domain string, event string, version string) error
}

type Topic struct {
	Domain  string
	Event   string
	Version string

	valid Validator
}

func NewTopic(domain string, event string, version string, valid Validator) *Topic {
	return &Topic{
		Domain:  domain,
		Event:   event,
		Version: version,
		valid:   valid,
	}
}

func (t *Topic) GetName() string {
	return fmt.Sprintf("%s.%s.v%s", t.Domain, t.Event, t.Version)
}

func (t *Topic) ValidateBytes(subject []byte) error {
	return t.valid.ValidateBytes(subject, t.Domain, t.Event, t.Version)
}

func (t *Topic) ValidateString(subject string) error {
	return t.valid.ValidateString(subject, t.Domain, t.Event, t.Version)
}
