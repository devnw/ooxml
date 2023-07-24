package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HeadersConstructor(t *testing.T) {
	v := wml.NewCT_Headers()
	if v == nil {
		t.Errorf("wml.NewCT_Headers must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Headers should validate: %s", err)
	}
}

func TestCT_HeadersMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Headers()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Headers()
	xml.Unmarshal(buf, v2)
}
