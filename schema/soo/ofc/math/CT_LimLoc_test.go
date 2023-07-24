package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_LimLocConstructor(t *testing.T) {
	v := math.NewCT_LimLoc()
	if v == nil {
		t.Errorf("math.NewCT_LimLoc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_LimLoc should validate: %s", err)
	}
}

func TestCT_LimLocMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_LimLoc()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_LimLoc()
	xml.Unmarshal(buf, v2)
}
