package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_EqArrPrConstructor(t *testing.T) {
	v := math.NewCT_EqArrPr()
	if v == nil {
		t.Errorf("math.NewCT_EqArrPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_EqArrPr should validate: %s", err)
	}
}

func TestCT_EqArrPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_EqArrPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_EqArrPr()
	xml.Unmarshal(buf, v2)
}
