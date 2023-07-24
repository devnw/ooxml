package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TabsConstructor(t *testing.T) {
	v := wml.NewCT_Tabs()
	if v == nil {
		t.Errorf("wml.NewCT_Tabs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Tabs should validate: %s", err)
	}
}

func TestCT_TabsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Tabs()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Tabs()
	xml.Unmarshal(buf, v2)
}
