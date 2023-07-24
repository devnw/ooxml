package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualContentPartPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualContentPartProperties()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualContentPartProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualContentPartProperties should validate: %s", err)
	}
}

func TestCT_NonVisualContentPartPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualContentPartProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualContentPartProperties()
	xml.Unmarshal(buf, v2)
}
