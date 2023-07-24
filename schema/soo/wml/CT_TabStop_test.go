package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TabStopConstructor(t *testing.T) {
	v := wml.NewCT_TabStop()
	if v == nil {
		t.Errorf("wml.NewCT_TabStop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TabStop should validate: %s", err)
	}
}

func TestCT_TabStopMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TabStop()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TabStop()
	xml.Unmarshal(buf, v2)
}
