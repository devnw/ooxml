package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_StyleConstructor(t *testing.T) {
	v := wml.NewCT_Style()
	if v == nil {
		t.Errorf("wml.NewCT_Style must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Style should validate: %s", err)
	}
}

func TestCT_StyleMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Style()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Style()
	xml.Unmarshal(buf, v2)
}
