package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_Integer255Constructor(t *testing.T) {
	v := math.NewCT_Integer255()
	if v == nil {
		t.Errorf("math.NewCT_Integer255 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Integer255 should validate: %s", err)
	}
}

func TestCT_Integer255MarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Integer255()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Integer255()
	xml.Unmarshal(buf, v2)
}
