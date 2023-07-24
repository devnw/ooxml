package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocVarConstructor(t *testing.T) {
	v := wml.NewCT_DocVar()
	if v == nil {
		t.Errorf("wml.NewCT_DocVar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocVar should validate: %s", err)
	}
}

func TestCT_DocVarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocVar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocVar()
	xml.Unmarshal(buf, v2)
}
