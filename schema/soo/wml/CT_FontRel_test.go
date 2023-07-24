package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontRelConstructor(t *testing.T) {
	v := wml.NewCT_FontRel()
	if v == nil {
		t.Errorf("wml.NewCT_FontRel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FontRel should validate: %s", err)
	}
}

func TestCT_FontRelMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FontRel()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FontRel()
	xml.Unmarshal(buf, v2)
}
