package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartConstructor(t *testing.T) {
	v := wml.NewCT_DocPart()
	if v == nil {
		t.Errorf("wml.NewCT_DocPart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPart should validate: %s", err)
	}
}

func TestCT_DocPartMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPart()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPart()
	xml.Unmarshal(buf, v2)
}
