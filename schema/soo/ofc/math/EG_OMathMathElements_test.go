package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestEG_OMathMathElementsConstructor(t *testing.T) {
	v := math.NewEG_OMathMathElements()
	if v == nil {
		t.Errorf("math.NewEG_OMathMathElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_OMathMathElements should validate: %s", err)
	}
}

func TestEG_OMathMathElementsMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_OMathMathElements()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_OMathMathElements()
	xml.Unmarshal(buf, v2)
}
