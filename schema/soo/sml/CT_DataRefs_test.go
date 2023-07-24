package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataRefsConstructor(t *testing.T) {
	v := sml.NewCT_DataRefs()
	if v == nil {
		t.Errorf("sml.NewCT_DataRefs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataRefs should validate: %s", err)
	}
}

func TestCT_DataRefsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataRefs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataRefs()
	xml.Unmarshal(buf, v2)
}
