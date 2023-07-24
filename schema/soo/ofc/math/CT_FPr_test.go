package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_FPrConstructor(t *testing.T) {
	v := math.NewCT_FPr()
	if v == nil {
		t.Errorf("math.NewCT_FPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_FPr should validate: %s", err)
	}
}

func TestCT_FPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_FPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_FPr()
	xml.Unmarshal(buf, v2)
}
