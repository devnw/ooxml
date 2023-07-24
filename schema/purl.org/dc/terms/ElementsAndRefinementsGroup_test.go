package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestElementsAndRefinementsGroupConstructor(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroup()
	if v == nil {
		t.Errorf("terms.NewElementsAndRefinementsGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementsAndRefinementsGroup should validate: %s", err)
	}
}

func TestElementsAndRefinementsGroupMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroup()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementsAndRefinementsGroup()
	xml.Unmarshal(buf, v2)
}
