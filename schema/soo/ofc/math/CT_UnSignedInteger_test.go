package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_UnSignedIntegerConstructor(t *testing.T) {
	v := math.NewCT_UnSignedInteger()
	if v == nil {
		t.Errorf("math.NewCT_UnSignedInteger must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_UnSignedInteger should validate: %s", err)
	}
}

func TestCT_UnSignedIntegerMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_UnSignedInteger()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_UnSignedInteger()
	xml.Unmarshal(buf, v2)
}
