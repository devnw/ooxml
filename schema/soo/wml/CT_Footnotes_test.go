package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FootnotesConstructor(t *testing.T) {
	v := wml.NewCT_Footnotes()
	if v == nil {
		t.Errorf("wml.NewCT_Footnotes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Footnotes should validate: %s", err)
	}
}

func TestCT_FootnotesMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Footnotes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Footnotes()
	xml.Unmarshal(buf, v2)
}
