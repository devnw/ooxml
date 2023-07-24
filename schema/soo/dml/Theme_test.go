package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestThemeConstructor(t *testing.T) {
	v := dml.NewTheme()
	if v == nil {
		t.Errorf("dml.NewTheme must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Theme should validate: %s", err)
	}
}

func TestThemeMarshalUnmarshal(t *testing.T) {
	v := dml.NewTheme()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewTheme()
	xml.Unmarshal(buf, v2)
}
