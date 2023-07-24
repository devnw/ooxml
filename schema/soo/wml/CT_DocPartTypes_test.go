package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartTypesConstructor(t *testing.T) {
	v := wml.NewCT_DocPartTypes()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartTypes should validate: %s", err)
	}
}

func TestCT_DocPartTypesMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartTypes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartTypes()
	xml.Unmarshal(buf, v2)
}
