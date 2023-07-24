package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalcPrConstructor(t *testing.T) {
	v := sml.NewCT_CalcPr()
	if v == nil {
		t.Errorf("sml.NewCT_CalcPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalcPr should validate: %s", err)
	}
}

func TestCT_CalcPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalcPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalcPr()
	xml.Unmarshal(buf, v2)
}
