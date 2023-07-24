package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocumentBaseConstructor(t *testing.T) {
	v := wml.NewCT_DocumentBase()
	if v == nil {
		t.Errorf("wml.NewCT_DocumentBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocumentBase should validate: %s", err)
	}
}

func TestCT_DocumentBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocumentBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocumentBase()
	xml.Unmarshal(buf, v2)
}
