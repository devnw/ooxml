package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ProofConstructor(t *testing.T) {
	v := wml.NewCT_Proof()
	if v == nil {
		t.Errorf("wml.NewCT_Proof must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Proof should validate: %s", err)
	}
}

func TestCT_ProofMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Proof()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Proof()
	xml.Unmarshal(buf, v2)
}
