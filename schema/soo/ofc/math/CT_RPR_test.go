package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_RPRConstructor(t *testing.T) {
	v := math.NewCT_RPR()
	if v == nil {
		t.Errorf("math.NewCT_RPR must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_RPR should validate: %s", err)
	}
}

func TestCT_RPRMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_RPR()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_RPR()
	xml.Unmarshal(buf, v2)
}
