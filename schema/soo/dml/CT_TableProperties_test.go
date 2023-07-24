package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TablePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TableProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TableProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableProperties should validate: %s", err)
	}
}

func TestCT_TablePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableProperties()
	xml.Unmarshal(buf, v2)
}
