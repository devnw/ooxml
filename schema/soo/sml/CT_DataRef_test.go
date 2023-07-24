package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataRefConstructor(t *testing.T) {
	v := sml.NewCT_DataRef()
	if v == nil {
		t.Errorf("sml.NewCT_DataRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataRef should validate: %s", err)
	}
}

func TestCT_DataRefMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataRef()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataRef()
	xml.Unmarshal(buf, v2)
}
