package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestBlipConstructor(t *testing.T) {
	v := dml.NewBlip()
	if v == nil {
		t.Errorf("dml.NewBlip must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Blip should validate: %s", err)
	}
}

func TestBlipMarshalUnmarshal(t *testing.T) {
	v := dml.NewBlip()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewBlip()
	xml.Unmarshal(buf, v2)
}
