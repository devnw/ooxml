package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestElementsAndRefinementsGroupChoiceConstructor(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroupChoice()
	if v == nil {
		t.Errorf("terms.NewElementsAndRefinementsGroupChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementsAndRefinementsGroupChoice should validate: %s", err)
	}
}

func TestElementsAndRefinementsGroupChoiceMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroupChoice()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementsAndRefinementsGroupChoice()
	xml.Unmarshal(buf, v2)
}
