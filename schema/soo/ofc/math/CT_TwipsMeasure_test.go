package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_TwipsMeasureConstructor(t *testing.T) {
	v := math.NewCT_TwipsMeasure()
	if v == nil {
		t.Errorf("math.NewCT_TwipsMeasure must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_TwipsMeasure should validate: %s", err)
	}
}

func TestCT_TwipsMeasureMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_TwipsMeasure()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_TwipsMeasure()
	xml.Unmarshal(buf, v2)
}
