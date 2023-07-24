package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BoxPrConstructor(t *testing.T) {
	v := math.NewCT_BoxPr()
	if v == nil {
		t.Errorf("math.NewCT_BoxPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BoxPr should validate: %s", err)
	}
}

func TestCT_BoxPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BoxPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BoxPr()
	xml.Unmarshal(buf, v2)
}
