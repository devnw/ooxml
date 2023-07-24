package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LineNumberConstructor(t *testing.T) {
	v := wml.NewCT_LineNumber()
	if v == nil {
		t.Errorf("wml.NewCT_LineNumber must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LineNumber should validate: %s", err)
	}
}

func TestCT_LineNumberMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LineNumber()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LineNumber()
	xml.Unmarshal(buf, v2)
}
