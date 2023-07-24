package elements_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/elements"
)

func TestElementsGroupChoiceConstructor(t *testing.T) {
	v := elements.NewElementsGroupChoice()
	if v == nil {
		t.Errorf("elements.NewElementsGroupChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.ElementsGroupChoice should validate: %s", err)
	}
}

func TestElementsGroupChoiceMarshalUnmarshal(t *testing.T) {
	v := elements.NewElementsGroupChoice()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewElementsGroupChoice()
	xml.Unmarshal(buf, v2)
}
