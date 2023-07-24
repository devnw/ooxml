package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_IndConstructor(t *testing.T) {
	v := wml.NewCT_Ind()
	if v == nil {
		t.Errorf("wml.NewCT_Ind must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Ind should validate: %s", err)
	}
}

func TestCT_IndMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Ind()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Ind()
	xml.Unmarshal(buf, v2)
}
