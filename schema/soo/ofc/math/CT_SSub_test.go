package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSubConstructor(t *testing.T) {
	v := math.NewCT_SSub()
	if v == nil {
		t.Errorf("math.NewCT_SSub must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSub should validate: %s", err)
	}
}

func TestCT_SSubMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSub()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSub()
	xml.Unmarshal(buf, v2)
}
