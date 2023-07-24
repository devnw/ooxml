package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CharsetConstructor(t *testing.T) {
	v := wml.NewCT_Charset()
	if v == nil {
		t.Errorf("wml.NewCT_Charset must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Charset should validate: %s", err)
	}
}

func TestCT_CharsetMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Charset()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Charset()
	xml.Unmarshal(buf, v2)
}
