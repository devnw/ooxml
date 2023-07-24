package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RPrOriginalConstructor(t *testing.T) {
	v := wml.NewCT_RPrOriginal()
	if v == nil {
		t.Errorf("wml.NewCT_RPrOriginal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RPrOriginal should validate: %s", err)
	}
}

func TestCT_RPrOriginalMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RPrOriginal()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RPrOriginal()
	xml.Unmarshal(buf, v2)
}
