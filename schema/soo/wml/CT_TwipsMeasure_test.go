package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TwipsMeasureConstructor(t *testing.T) {
	v := wml.NewCT_TwipsMeasure()
	if v == nil {
		t.Errorf("wml.NewCT_TwipsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TwipsMeasure should validate: %s", err)
	}
}

func TestCT_TwipsMeasureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TwipsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TwipsMeasure()
	xml.Unmarshal(buf, v2)
}
