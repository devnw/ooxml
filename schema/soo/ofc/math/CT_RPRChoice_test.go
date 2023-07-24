package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_RPRChoiceConstructor(t *testing.T) {
	v := math.NewCT_RPRChoice()
	if v == nil {
		t.Errorf("math.NewCT_RPRChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_RPRChoice should validate: %s", err)
	}
}

func TestCT_RPRChoiceMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_RPRChoice()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_RPRChoice()
	xml.Unmarshal(buf, v2)
}
