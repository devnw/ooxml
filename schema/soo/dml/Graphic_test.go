package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestGraphicConstructor(t *testing.T) {
	v := dml.NewGraphic()
	if v == nil {
		t.Errorf("dml.NewGraphic must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Graphic should validate: %s", err)
	}
}

func TestGraphicMarshalUnmarshal(t *testing.T) {
	v := dml.NewGraphic()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewGraphic()
	xml.Unmarshal(buf, v2)
}
