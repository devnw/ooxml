package elements_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/elements"
)

func TestAnyConstructor(t *testing.T) {
	v := elements.NewAny()
	if v == nil {
		t.Errorf("elements.NewAny must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.Any should validate: %s", err)
	}
}

func TestAnyMarshalUnmarshal(t *testing.T) {
	v := elements.NewAny()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewAny()
	xml.Unmarshal(buf, v2)
}
