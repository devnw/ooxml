package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocVarsConstructor(t *testing.T) {
	v := wml.NewCT_DocVars()
	if v == nil {
		t.Errorf("wml.NewCT_DocVars must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocVars should validate: %s", err)
	}
}

func TestCT_DocVarsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocVars()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocVars()
	xml.Unmarshal(buf, v2)
}
