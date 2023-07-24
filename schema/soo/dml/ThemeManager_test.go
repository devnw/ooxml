package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestThemeManagerConstructor(t *testing.T) {
	v := dml.NewThemeManager()
	if v == nil {
		t.Errorf("dml.NewThemeManager must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.ThemeManager should validate: %s", err)
	}
}

func TestThemeManagerMarshalUnmarshal(t *testing.T) {
	v := dml.NewThemeManager()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewThemeManager()
	xml.Unmarshal(buf, v2)
}
