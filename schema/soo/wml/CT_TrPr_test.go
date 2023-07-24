package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TrPrConstructor(t *testing.T) {
	v := wml.NewCT_TrPr()
	if v == nil {
		t.Errorf("wml.NewCT_TrPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrPr should validate: %s", err)
	}
}

func TestCT_TrPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrPr()
	xml.Unmarshal(buf, v2)
}
