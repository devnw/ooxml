package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SSubSupConstructor(t *testing.T) {
	v := math.NewCT_SSubSup()
	if v == nil {
		t.Errorf("math.NewCT_SSubSup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSubSup should validate: %s", err)
	}
}

func TestCT_SSubSupMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSubSup()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSubSup()
	xml.Unmarshal(buf, v2)
}
