package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TextboxTightWrapConstructor(t *testing.T) {
	v := wml.NewCT_TextboxTightWrap()
	if v == nil {
		t.Errorf("wml.NewCT_TextboxTightWrap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextboxTightWrap should validate: %s", err)
	}
}

func TestCT_TextboxTightWrapMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextboxTightWrap()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextboxTightWrap()
	xml.Unmarshal(buf, v2)
}
