package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartTypeConstructor(t *testing.T) {
	v := wml.NewCT_DocPartType()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartType should validate: %s", err)
	}
}

func TestCT_DocPartTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartType()
	xml.Unmarshal(buf, v2)
}
