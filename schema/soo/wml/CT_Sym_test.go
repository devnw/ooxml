package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SymConstructor(t *testing.T) {
	v := wml.NewCT_Sym()
	if v == nil {
		t.Errorf("wml.NewCT_Sym must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Sym should validate: %s", err)
	}
}

func TestCT_SymMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Sym()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Sym()
	xml.Unmarshal(buf, v2)
}
