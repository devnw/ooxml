package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GeomRectConstructor(t *testing.T) {
	v := dml.NewCT_GeomRect()
	if v == nil {
		t.Errorf("dml.NewCT_GeomRect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GeomRect should validate: %s", err)
	}
}

func TestCT_GeomRectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GeomRect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GeomRect()
	xml.Unmarshal(buf, v2)
}
