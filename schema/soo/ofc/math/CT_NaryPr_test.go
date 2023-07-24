package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_NaryPrConstructor(t *testing.T) {
	v := math.NewCT_NaryPr()
	if v == nil {
		t.Errorf("math.NewCT_NaryPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_NaryPr should validate: %s", err)
	}
}

func TestCT_NaryPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_NaryPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_NaryPr()
	xml.Unmarshal(buf, v2)
}
