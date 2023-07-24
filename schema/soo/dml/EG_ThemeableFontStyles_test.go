package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_ThemeableFontStylesConstructor(t *testing.T) {
	v := dml.NewEG_ThemeableFontStyles()
	if v == nil {
		t.Errorf("dml.NewEG_ThemeableFontStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ThemeableFontStyles should validate: %s", err)
	}
}

func TestEG_ThemeableFontStylesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ThemeableFontStyles()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ThemeableFontStyles()
	xml.Unmarshal(buf, v2)
}
