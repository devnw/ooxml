package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PageMarConstructor(t *testing.T) {
	v := wml.NewCT_PageMar()
	if v == nil {
		t.Errorf("wml.NewCT_PageMar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageMar should validate: %s", err)
	}
}

func TestCT_PageMarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageMar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageMar()
	xml.Unmarshal(buf, v2)
}
