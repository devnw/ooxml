package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartsConstructor(t *testing.T) {
	v := wml.NewCT_DocParts()
	if v == nil {
		t.Errorf("wml.NewCT_DocParts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocParts should validate: %s", err)
	}
}

func TestCT_DocPartsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocParts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocParts()
	xml.Unmarshal(buf, v2)
}
