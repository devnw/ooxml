package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_PhantConstructor(t *testing.T) {
	v := math.NewCT_Phant()
	if v == nil {
		t.Errorf("math.NewCT_Phant must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Phant should validate: %s", err)
	}
}

func TestCT_PhantMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Phant()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Phant()
	xml.Unmarshal(buf, v2)
}
