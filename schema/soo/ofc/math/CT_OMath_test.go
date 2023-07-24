package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathConstructor(t *testing.T) {
	v := math.NewCT_OMath()
	if v == nil {
		t.Errorf("math.NewCT_OMath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMath should validate: %s", err)
	}
}

func TestCT_OMathMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMath()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMath()
	xml.Unmarshal(buf, v2)
}
