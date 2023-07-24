package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_ThemeableEffectStyleConstructor(t *testing.T) {
	v := dml.NewEG_ThemeableEffectStyle()
	if v == nil {
		t.Errorf("dml.NewEG_ThemeableEffectStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ThemeableEffectStyle should validate: %s", err)
	}
}

func TestEG_ThemeableEffectStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ThemeableEffectStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ThemeableEffectStyle()
	xml.Unmarshal(buf, v2)
}
