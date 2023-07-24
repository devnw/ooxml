package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumFmtConstructor(t *testing.T) {
	v := wml.NewCT_NumFmt()
	if v == nil {
		t.Errorf("wml.NewCT_NumFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumFmt should validate: %s", err)
	}
}

func TestCT_NumFmtMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumFmt()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumFmt()
	xml.Unmarshal(buf, v2)
}
