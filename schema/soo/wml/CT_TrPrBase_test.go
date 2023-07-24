package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TrPrBaseConstructor(t *testing.T) {
	v := wml.NewCT_TrPrBase()
	if v == nil {
		t.Errorf("wml.NewCT_TrPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrPrBase should validate: %s", err)
	}
}

func TestCT_TrPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrPrBase()
	xml.Unmarshal(buf, v2)
}
