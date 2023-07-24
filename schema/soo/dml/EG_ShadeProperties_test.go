package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_ShadePropertiesConstructor(t *testing.T) {
	v := dml.NewEG_ShadeProperties()
	if v == nil {
		t.Errorf("dml.NewEG_ShadeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ShadeProperties should validate: %s", err)
	}
}

func TestEG_ShadePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ShadeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ShadeProperties()
	xml.Unmarshal(buf, v2)
}
