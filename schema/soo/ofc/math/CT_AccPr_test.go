package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_AccPrConstructor(t *testing.T) {
	v := math.NewCT_AccPr()
	if v == nil {
		t.Errorf("math.NewCT_AccPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_AccPr should validate: %s", err)
	}
}

func TestCT_AccPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_AccPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_AccPr()
	xml.Unmarshal(buf, v2)
}
