package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathArgConstructor(t *testing.T) {
	v := math.NewCT_OMathArg()
	if v == nil {
		t.Errorf("math.NewCT_OMathArg must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMathArg should validate: %s", err)
	}
}

func TestCT_OMathArgMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMathArg()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMathArg()
	xml.Unmarshal(buf, v2)
}
