package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestMathPrConstructor(t *testing.T) {
	v := math.NewMathPr()
	if v == nil {
		t.Errorf("math.NewMathPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.MathPr should validate: %s", err)
	}
}

func TestMathPrMarshalUnmarshal(t *testing.T) {
	v := math.NewMathPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewMathPr()
	xml.Unmarshal(buf, v2)
}
