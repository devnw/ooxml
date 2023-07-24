package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ModifyVerifierConstructor(t *testing.T) {
	v := pml.NewCT_ModifyVerifier()
	if v == nil {
		t.Errorf("pml.NewCT_ModifyVerifier must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ModifyVerifier should validate: %s", err)
	}
}

func TestCT_ModifyVerifierMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ModifyVerifier()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ModifyVerifier()
	xml.Unmarshal(buf, v2)
}
