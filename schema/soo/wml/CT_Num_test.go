package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumConstructor(t *testing.T) {
	v := wml.NewCT_Num()
	if v == nil {
		t.Errorf("wml.NewCT_Num must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Num should validate: %s", err)
	}
}

func TestCT_NumMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Num()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Num()
	xml.Unmarshal(buf, v2)
}
