package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BaseStylesOverrideConstructor(t *testing.T) {
	v := dml.NewCT_BaseStylesOverride()
	if v == nil {
		t.Errorf("dml.NewCT_BaseStylesOverride must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BaseStylesOverride should validate: %s", err)
	}
}

func TestCT_BaseStylesOverrideMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BaseStylesOverride()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BaseStylesOverride()
	xml.Unmarshal(buf, v2)
}
