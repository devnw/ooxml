package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_GroupChrPrConstructor(t *testing.T) {
	v := math.NewCT_GroupChrPr()
	if v == nil {
		t.Errorf("math.NewCT_GroupChrPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_GroupChrPr should validate: %s", err)
	}
}

func TestCT_GroupChrPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_GroupChrPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_GroupChrPr()
	xml.Unmarshal(buf, v2)
}
