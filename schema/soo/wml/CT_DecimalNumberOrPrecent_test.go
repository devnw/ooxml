package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DecimalNumberOrPrecentConstructor(t *testing.T) {
	v := wml.NewCT_DecimalNumberOrPrecent()
	if v == nil {
		t.Errorf("wml.NewCT_DecimalNumberOrPrecent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DecimalNumberOrPrecent should validate: %s", err)
	}
}

func TestCT_DecimalNumberOrPrecentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DecimalNumberOrPrecent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DecimalNumberOrPrecent()
	xml.Unmarshal(buf, v2)
}
