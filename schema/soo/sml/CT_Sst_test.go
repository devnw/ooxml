package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SstConstructor(t *testing.T) {
	v := sml.NewCT_Sst()
	if v == nil {
		t.Errorf("sml.NewCT_Sst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Sst should validate: %s", err)
	}
}

func TestCT_SstMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Sst()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Sst()
	xml.Unmarshal(buf, v2)
}
