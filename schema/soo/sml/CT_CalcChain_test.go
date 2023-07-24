package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalcChainConstructor(t *testing.T) {
	v := sml.NewCT_CalcChain()
	if v == nil {
		t.Errorf("sml.NewCT_CalcChain must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalcChain should validate: %s", err)
	}
}

func TestCT_CalcChainMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalcChain()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalcChain()
	xml.Unmarshal(buf, v2)
}
