package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_MCSConstructor(t *testing.T) {
	v := math.NewCT_MCS()
	if v == nil {
		t.Errorf("math.NewCT_MCS must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MCS should validate: %s", err)
	}
}

func TestCT_MCSMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MCS()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MCS()
	xml.Unmarshal(buf, v2)
}
