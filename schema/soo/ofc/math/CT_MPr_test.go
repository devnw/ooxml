package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_MPrConstructor(t *testing.T) {
	v := math.NewCT_MPr()
	if v == nil {
		t.Errorf("math.NewCT_MPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MPr should validate: %s", err)
	}
}

func TestCT_MPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MPr()
	xml.Unmarshal(buf, v2)
}
