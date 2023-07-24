package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SPrePrConstructor(t *testing.T) {
	v := math.NewCT_SPrePr()
	if v == nil {
		t.Errorf("math.NewCT_SPrePr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SPrePr should validate: %s", err)
	}
}

func TestCT_SPrePrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SPrePr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SPrePr()
	xml.Unmarshal(buf, v2)
}
