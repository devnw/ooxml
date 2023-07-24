package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_ManualBreakConstructor(t *testing.T) {
	v := math.NewCT_ManualBreak()
	if v == nil {
		t.Errorf("math.NewCT_ManualBreak must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_ManualBreak should validate: %s", err)
	}
}

func TestCT_ManualBreakMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_ManualBreak()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_ManualBreak()
	xml.Unmarshal(buf, v2)
}
