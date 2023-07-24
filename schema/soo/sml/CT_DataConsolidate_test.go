package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataConsolidateConstructor(t *testing.T) {
	v := sml.NewCT_DataConsolidate()
	if v == nil {
		t.Errorf("sml.NewCT_DataConsolidate must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataConsolidate should validate: %s", err)
	}
}

func TestCT_DataConsolidateMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataConsolidate()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataConsolidate()
	xml.Unmarshal(buf, v2)
}
