package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LongHexNumberConstructor(t *testing.T) {
	v := wml.NewCT_LongHexNumber()
	if v == nil {
		t.Errorf("wml.NewCT_LongHexNumber must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LongHexNumber should validate: %s", err)
	}
}

func TestCT_LongHexNumberMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LongHexNumber()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LongHexNumber()
	xml.Unmarshal(buf, v2)
}
