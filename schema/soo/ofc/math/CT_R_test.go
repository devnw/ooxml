package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_RConstructor(t *testing.T) {
	v := math.NewCT_R()
	if v == nil {
		t.Errorf("math.NewCT_R must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_R should validate: %s", err)
	}
}

func TestCT_RMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_R()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_R()
	xml.Unmarshal(buf, v2)
}
