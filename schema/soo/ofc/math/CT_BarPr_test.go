package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BarPrConstructor(t *testing.T) {
	v := math.NewCT_BarPr()
	if v == nil {
		t.Errorf("math.NewCT_BarPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BarPr should validate: %s", err)
	}
}

func TestCT_BarPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BarPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BarPr()
	xml.Unmarshal(buf, v2)
}
