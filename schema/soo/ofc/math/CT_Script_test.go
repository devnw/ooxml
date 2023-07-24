package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_ScriptConstructor(t *testing.T) {
	v := math.NewCT_Script()
	if v == nil {
		t.Errorf("math.NewCT_Script must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Script should validate: %s", err)
	}
}

func TestCT_ScriptMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Script()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Script()
	xml.Unmarshal(buf, v2)
}
