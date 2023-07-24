package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocumentConstructor(t *testing.T) {
	v := wml.NewCT_Document()
	if v == nil {
		t.Errorf("wml.NewCT_Document must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Document should validate: %s", err)
	}
}

func TestCT_DocumentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Document()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Document()
	xml.Unmarshal(buf, v2)
}
