package relationships_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/relationships"
)

func TestRelationshipConstructor(t *testing.T) {
	v := relationships.NewRelationship()
	if v == nil {
		t.Errorf("relationships.NewRelationship must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed relationships.Relationship should validate: %s", err)
	}
}

func TestRelationshipMarshalUnmarshal(t *testing.T) {
	v := relationships.NewRelationship()
	buf, _ := xml.Marshal(v)
	v2 := relationships.NewRelationship()
	xml.Unmarshal(buf, v2)
}
