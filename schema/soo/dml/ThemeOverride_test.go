package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestThemeOverrideConstructor(t *testing.T) {
	v := dml.NewThemeOverride()
	if v == nil {
		t.Errorf("dml.NewThemeOverride must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.ThemeOverride should validate: %s", err)
	}
}

func TestThemeOverrideMarshalUnmarshal(t *testing.T) {
	v := dml.NewThemeOverride()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewThemeOverride()
	xml.Unmarshal(buf, v2)
}
