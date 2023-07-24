package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SmartTagPrConstructor(t *testing.T) {
	v := wml.NewCT_SmartTagPr()
	if v == nil {
		t.Errorf("wml.NewCT_SmartTagPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SmartTagPr should validate: %s", err)
	}
}

func TestCT_SmartTagPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SmartTagPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SmartTagPr()
	xml.Unmarshal(buf, v2)
}
