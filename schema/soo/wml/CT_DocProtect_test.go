package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocProtectConstructor(t *testing.T) {
	v := wml.NewCT_DocProtect()
	if v == nil {
		t.Errorf("wml.NewCT_DocProtect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocProtect should validate: %s", err)
	}
}

func TestCT_DocProtectMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocProtect()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocProtect()
	xml.Unmarshal(buf, v2)
}
