package relationships_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/relationships"
)

func TestCT_RelationshipConstructor(t *testing.T) {
	v := relationships.NewCT_Relationship()
	if v == nil {
		t.Errorf("relationships.NewCT_Relationship must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed relationships.CT_Relationship should validate: %s", err)
	}
}

func TestCT_RelationshipMarshalUnmarshal(t *testing.T) {
	v := relationships.NewCT_Relationship()
	buf, _ := xml.Marshal(v)
	v2 := relationships.NewCT_Relationship()
	xml.Unmarshal(buf, v2)
}
