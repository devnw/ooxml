package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PageSzConstructor(t *testing.T) {
	v := wml.NewCT_PageSz()
	if v == nil {
		t.Errorf("wml.NewCT_PageSz must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageSz should validate: %s", err)
	}
}

func TestCT_PageSzMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageSz()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageSz()
	xml.Unmarshal(buf, v2)
}
