package elements_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/elements"
)

func TestSimpleLiteralConstructor(t *testing.T) {
	v := elements.NewSimpleLiteral()
	if v == nil {
		t.Errorf("elements.NewSimpleLiteral must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.SimpleLiteral should validate: %s", err)
	}
}

func TestSimpleLiteralMarshalUnmarshal(t *testing.T) {
	v := elements.NewSimpleLiteral()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewSimpleLiteral()
	xml.Unmarshal(buf, v2)
}
