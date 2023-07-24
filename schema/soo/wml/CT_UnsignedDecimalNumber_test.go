package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_UnsignedDecimalNumberConstructor(t *testing.T) {
	v := wml.NewCT_UnsignedDecimalNumber()
	if v == nil {
		t.Errorf("wml.NewCT_UnsignedDecimalNumber must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_UnsignedDecimalNumber should validate: %s", err)
	}
}

func TestCT_UnsignedDecimalNumberMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_UnsignedDecimalNumber()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_UnsignedDecimalNumber()
	xml.Unmarshal(buf, v2)
}
