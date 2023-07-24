package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_ShpConstructor(t *testing.T) {
	v := math.NewCT_Shp()
	if v == nil {
		t.Errorf("math.NewCT_Shp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Shp should validate: %s", err)
	}
}

func TestCT_ShpMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Shp()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Shp()
	xml.Unmarshal(buf, v2)
}
