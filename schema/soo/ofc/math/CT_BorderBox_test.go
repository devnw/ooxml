package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_BorderBoxConstructor(t *testing.T) {
	v := math.NewCT_BorderBox()
	if v == nil {
		t.Errorf("math.NewCT_BorderBox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_BorderBox should validate: %s", err)
	}
}

func TestCT_BorderBoxMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_BorderBox()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_BorderBox()
	xml.Unmarshal(buf, v2)
}
