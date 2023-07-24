package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EndnotesConstructor(t *testing.T) {
	v := wml.NewCT_Endnotes()
	if v == nil {
		t.Errorf("wml.NewCT_Endnotes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Endnotes should validate: %s", err)
	}
}

func TestCT_EndnotesMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Endnotes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Endnotes()
	xml.Unmarshal(buf, v2)
}
