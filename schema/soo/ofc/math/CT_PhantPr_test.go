package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_PhantPrConstructor(t *testing.T) {
	v := math.NewCT_PhantPr()
	if v == nil {
		t.Errorf("math.NewCT_PhantPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_PhantPr should validate: %s", err)
	}
}

func TestCT_PhantPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_PhantPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_PhantPr()
	xml.Unmarshal(buf, v2)
}
