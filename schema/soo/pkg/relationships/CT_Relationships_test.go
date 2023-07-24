package relationships_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/relationships"
)

func TestCT_RelationshipsConstructor(t *testing.T) {
	v := relationships.NewCT_Relationships()
	if v == nil {
		t.Errorf("relationships.NewCT_Relationships must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed relationships.CT_Relationships should validate: %s", err)
	}
}

func TestCT_RelationshipsMarshalUnmarshal(t *testing.T) {
	v := relationships.NewCT_Relationships()
	buf, _ := xml.Marshal(v)
	v2 := relationships.NewCT_Relationships()
	xml.Unmarshal(buf, v2)
}
