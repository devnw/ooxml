package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestFootnotesConstructor(t *testing.T) {
	v := wml.NewFootnotes()
	if v == nil {
		t.Errorf("wml.NewFootnotes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Footnotes should validate: %s", err)
	}
}

func TestFootnotesMarshalUnmarshal(t *testing.T) {
	v := wml.NewFootnotes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewFootnotes()
	xml.Unmarshal(buf, v2)
}
