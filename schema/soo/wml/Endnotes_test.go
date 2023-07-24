package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEndnotesConstructor(t *testing.T) {
	v := wml.NewEndnotes()
	if v == nil {
		t.Errorf("wml.NewEndnotes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Endnotes should validate: %s", err)
	}
}

func TestEndnotesMarshalUnmarshal(t *testing.T) {
	v := wml.NewEndnotes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEndnotes()
	xml.Unmarshal(buf, v2)
}
