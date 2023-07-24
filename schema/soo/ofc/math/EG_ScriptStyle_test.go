package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestEG_ScriptStyleConstructor(t *testing.T) {
	v := math.NewEG_ScriptStyle()
	if v == nil {
		t.Errorf("math.NewEG_ScriptStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_ScriptStyle should validate: %s", err)
	}
}

func TestEG_ScriptStyleMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_ScriptStyle()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_ScriptStyle()
	xml.Unmarshal(buf, v2)
}
