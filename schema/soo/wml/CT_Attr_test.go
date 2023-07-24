package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AttrConstructor(t *testing.T) {
	v := wml.NewCT_Attr()
	if v == nil {
		t.Errorf("wml.NewCT_Attr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Attr should validate: %s", err)
	}
}

func TestCT_AttrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Attr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Attr()
	xml.Unmarshal(buf, v2)
}
