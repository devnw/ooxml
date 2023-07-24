package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BreakBinConstructor(t *testing.T) {
	v := math.NewCT_BreakBin()
	if v == nil {
		t.Errorf("math.NewCT_BreakBin must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BreakBin should validate: %s", err)
	}
}

func TestCT_BreakBinMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BreakBin()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BreakBin()
	xml.Unmarshal(buf, v2)
}
