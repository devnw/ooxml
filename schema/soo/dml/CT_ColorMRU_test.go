package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorMRUConstructor(t *testing.T) {
	v := dml.NewCT_ColorMRU()
	if v == nil {
		t.Errorf("dml.NewCT_ColorMRU must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorMRU should validate: %s", err)
	}
}

func TestCT_ColorMRUMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorMRU()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorMRU()
	xml.Unmarshal(buf, v2)
}
