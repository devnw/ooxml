package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PageNumberConstructor(t *testing.T) {
	v := wml.NewCT_PageNumber()
	if v == nil {
		t.Errorf("wml.NewCT_PageNumber must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageNumber should validate: %s", err)
	}
}

func TestCT_PageNumberMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageNumber()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageNumber()
	xml.Unmarshal(buf, v2)
}
