package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_LimLowPrConstructor(t *testing.T) {
	v := math.NewCT_LimLowPr()
	if v == nil {
		t.Errorf("math.NewCT_LimLowPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_LimLowPr should validate: %s", err)
	}
}

func TestCT_LimLowPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_LimLowPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_LimLowPr()
	xml.Unmarshal(buf, v2)
}
