package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_RChoiceConstructor(t *testing.T) {
	v := math.NewCT_RChoice()
	if v == nil {
		t.Errorf("math.NewCT_RChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_RChoice should validate: %s", err)
	}
}

func TestCT_RChoiceMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_RChoice()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_RChoice()
	xml.Unmarshal(buf, v2)
}
