package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BorderConstructor(t *testing.T) {
	v := wml.NewCT_Border()
	if v == nil {
		t.Errorf("wml.NewCT_Border must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Border should validate: %s", err)
	}
}

func TestCT_BorderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Border()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Border()
	xml.Unmarshal(buf, v2)
}
