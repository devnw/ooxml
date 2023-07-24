package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestEG_OMathElementsConstructor(t *testing.T) {
	v := math.NewEG_OMathElements()
	if v == nil {
		t.Errorf("math.NewEG_OMathElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_OMathElements should validate: %s", err)
	}
}

func TestEG_OMathElementsMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_OMathElements()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_OMathElements()
	xml.Unmarshal(buf, v2)
}
