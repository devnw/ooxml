package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_MConstructor(t *testing.T) {
	v := math.NewCT_M()
	if v == nil {
		t.Errorf("math.NewCT_M must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_M should validate: %s", err)
	}
}

func TestCT_MMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_M()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_M()
	xml.Unmarshal(buf, v2)
}
