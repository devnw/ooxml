package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_TopBotConstructor(t *testing.T) {
	v := math.NewCT_TopBot()
	if v == nil {
		t.Errorf("math.NewCT_TopBot must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_TopBot should validate: %s", err)
	}
}

func TestCT_TopBotMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_TopBot()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_TopBot()
	xml.Unmarshal(buf, v2)
}
