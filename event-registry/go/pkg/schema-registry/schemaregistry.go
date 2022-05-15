package schemaregistry

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type EventSchemaRegistry struct {
	path    string
	schemas map[string]gojsonschema.JSONLoader
}

func NewRegistry(path string) *EventSchemaRegistry {
	return &EventSchemaRegistry{
		path:    path,
		schemas: make(map[string]gojsonschema.JSONLoader),
	}
}

// RegisterSchema
//	domain.event.version
//	domain/event_name/version.json
func (r *EventSchemaRegistry) RegisterSchema(domain string, event string, version string) {
	schema := fmt.Sprintf("%s.%s.v%s", domain, event, version)
	r.schemas[schema] = gojsonschema.NewReferenceLoader(
		fmt.Sprintf("file:///%s/%s/%s/%s.json",
			r.path, domain, event, version))
}

func (r *EventSchemaRegistry) Validate(subject string, domain string, event string, version string) error {
	schema := fmt.Sprintf("%s.%s.v%s", domain, event, version)
	subjectDoc := gojsonschema.NewStringLoader(subject)
	schemaDoc := r.schemas[schema]

	result, err := gojsonschema.Validate(schemaDoc, subjectDoc)
	if err != nil {
		return err
	}

	if !result.Valid() {
		errMsg := ""
		for _, desc := range result.Errors() {
			errMsg = fmt.Sprintf("%s- %s\n", errMsg, desc)
		}
		return fmt.Errorf("the document is not valid. see errors :\n%s", errMsg)
	}
	return nil
}

func (r *EventSchemaRegistry) TestCompatibility(subject string, domain string, event string, version string) error {
	return nil
}
