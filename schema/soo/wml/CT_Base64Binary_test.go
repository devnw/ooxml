package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_Base64BinaryConstructor(t *testing.T) {
	v := wml.NewCT_Base64Binary()
	if v == nil {
		t.Errorf("wml.NewCT_Base64Binary must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Base64Binary should validate: %s", err)
	}
}

func TestCT_Base64BinaryMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Base64Binary()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Base64Binary()
	xml.Unmarshal(buf, v2)
}
