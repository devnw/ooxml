package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PConstructor(t *testing.T) {
	v := wml.NewCT_P()
	if v == nil {
		t.Errorf("wml.NewCT_P must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_P should validate: %s", err)
	}
}

func TestCT_PMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_P()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_P()
	xml.Unmarshal(buf, v2)
}
