package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PixelsMeasureConstructor(t *testing.T) {
	v := wml.NewCT_PixelsMeasure()
	if v == nil {
		t.Errorf("wml.NewCT_PixelsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PixelsMeasure should validate: %s", err)
	}
}

func TestCT_PixelsMeasureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PixelsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PixelsMeasure()
	xml.Unmarshal(buf, v2)
}
