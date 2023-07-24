package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EmptyConstructor(t *testing.T) {
	v := wml.NewCT_Empty()
	if v == nil {
		t.Errorf("wml.NewCT_Empty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Empty should validate: %s", err)
	}
}

func TestCT_EmptyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Empty()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Empty()
	xml.Unmarshal(buf, v2)
}
