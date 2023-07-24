package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocTypeConstructor(t *testing.T) {
	v := wml.NewCT_DocType()
	if v == nil {
		t.Errorf("wml.NewCT_DocType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocType should validate: %s", err)
	}
}

func TestCT_DocTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocType()
	xml.Unmarshal(buf, v2)
}
