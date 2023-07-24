package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartPrConstructor(t *testing.T) {
	v := wml.NewCT_DocPartPr()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartPr should validate: %s", err)
	}
}

func TestCT_DocPartPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartPr()
	xml.Unmarshal(buf, v2)
}
