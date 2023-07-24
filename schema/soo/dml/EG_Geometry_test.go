package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_GeometryConstructor(t *testing.T) {
	v := dml.NewEG_Geometry()
	if v == nil {
		t.Errorf("dml.NewEG_Geometry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_Geometry should validate: %s", err)
	}
}

func TestEG_GeometryMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_Geometry()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_Geometry()
	xml.Unmarshal(buf, v2)
}
