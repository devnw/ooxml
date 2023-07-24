package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ProofErrConstructor(t *testing.T) {
	v := wml.NewCT_ProofErr()
	if v == nil {
		t.Errorf("wml.NewCT_ProofErr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ProofErr should validate: %s", err)
	}
}

func TestCT_ProofErrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ProofErr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ProofErr()
	xml.Unmarshal(buf, v2)
}
