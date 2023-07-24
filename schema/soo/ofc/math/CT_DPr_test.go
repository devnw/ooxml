package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_DPrConstructor(t *testing.T) {
	v := math.NewCT_DPr()
	if v == nil {
		t.Errorf("math.NewCT_DPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_DPr should validate: %s", err)
	}
}

func TestCT_DPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_DPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_DPr()
	xml.Unmarshal(buf, v2)
}
