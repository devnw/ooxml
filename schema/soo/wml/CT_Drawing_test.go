package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := wml.NewCT_Drawing()
	if v == nil {
		t.Errorf("wml.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
