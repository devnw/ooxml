package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_MCConstructor(t *testing.T) {
	v := math.NewCT_MC()
	if v == nil {
		t.Errorf("math.NewCT_MC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MC should validate: %s", err)
	}
}

func TestCT_MCMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MC()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MC()
	xml.Unmarshal(buf, v2)
}
