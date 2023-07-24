package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartNameConstructor(t *testing.T) {
	v := wml.NewCT_DocPartName()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartName should validate: %s", err)
	}
}

func TestCT_DocPartNameMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartName()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartName()
	xml.Unmarshal(buf, v2)
}
