package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BreakBinSubConstructor(t *testing.T) {
	v := math.NewCT_BreakBinSub()
	if v == nil {
		t.Errorf("math.NewCT_BreakBinSub must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BreakBinSub should validate: %s", err)
	}
}

func TestCT_BreakBinSubMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BreakBinSub()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BreakBinSub()
	xml.Unmarshal(buf, v2)
}
