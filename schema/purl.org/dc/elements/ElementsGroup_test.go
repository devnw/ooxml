package elements_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/elements"
)

func TestElementsGroupConstructor(t *testing.T) {
	v := elements.NewElementsGroup()
	if v == nil {
		t.Errorf("elements.NewElementsGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.ElementsGroup should validate: %s", err)
	}
}

func TestElementsGroupMarshalUnmarshal(t *testing.T) {
	v := elements.NewElementsGroup()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewElementsGroup()
	xml.Unmarshal(buf, v2)
}
