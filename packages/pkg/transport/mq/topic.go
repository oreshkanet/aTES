package mq

import "fmt"

type Topic struct {
	Domain  string
	Event   string
	Version string
}

func (t *Topic) GetName() string {
	return fmt.Sprintf("%s.%s.v%s", t.Domain, t.Event, t.Version)
}
