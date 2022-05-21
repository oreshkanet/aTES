package mq

import "fmt"

type Validator interface {
	ValidateBytes(subject []byte, domain string, event string, version string) error
}

type Topic struct {
	Domain  string
	Event   string
	Version string

	validator Validator
}

func (t *Topic) GetName() string {
	return fmt.Sprintf("%s.%s.v%s", t.Domain, t.Event, t.Version)
}
