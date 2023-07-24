package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MathCtrlInsConstructor(t *testing.T) {
	v := wml.NewCT_MathCtrlIns()
	if v == nil {
		t.Errorf("wml.NewCT_MathCtrlIns must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MathCtrlIns should validate: %s", err)
	}
}

func TestCT_MathCtrlInsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MathCtrlIns()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MathCtrlIns()
	xml.Unmarshal(buf, v2)
}
