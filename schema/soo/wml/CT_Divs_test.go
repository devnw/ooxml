package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DivsConstructor(t *testing.T) {
	v := wml.NewCT_Divs()
	if v == nil {
		t.Errorf("wml.NewCT_Divs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Divs should validate: %s", err)
	}
}

func TestCT_DivsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Divs()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Divs()
	xml.Unmarshal(buf, v2)
}
