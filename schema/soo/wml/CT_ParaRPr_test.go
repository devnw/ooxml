package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ParaRPrConstructor(t *testing.T) {
	v := wml.NewCT_ParaRPr()
	if v == nil {
		t.Errorf("wml.NewCT_ParaRPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ParaRPr should validate: %s", err)
	}
}

func TestCT_ParaRPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ParaRPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ParaRPr()
	xml.Unmarshal(buf, v2)
}
