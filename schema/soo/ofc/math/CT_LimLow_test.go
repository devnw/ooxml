package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_LimLowConstructor(t *testing.T) {
	v := math.NewCT_LimLow()
	if v == nil {
		t.Errorf("math.NewCT_LimLow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_LimLow should validate: %s", err)
	}
}

func TestCT_LimLowMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_LimLow()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_LimLow()
	xml.Unmarshal(buf, v2)
}
