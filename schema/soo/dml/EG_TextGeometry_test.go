package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextGeometryConstructor(t *testing.T) {
	v := dml.NewEG_TextGeometry()
	if v == nil {
		t.Errorf("dml.NewEG_TextGeometry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextGeometry should validate: %s", err)
	}
}

func TestEG_TextGeometryMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextGeometry()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextGeometry()
	xml.Unmarshal(buf, v2)
}
