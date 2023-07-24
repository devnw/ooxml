package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_EqArrConstructor(t *testing.T) {
	v := math.NewCT_EqArr()
	if v == nil {
		t.Errorf("math.NewCT_EqArr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_EqArr should validate: %s", err)
	}
}

func TestCT_EqArrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_EqArr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_EqArr()
	xml.Unmarshal(buf, v2)
}
