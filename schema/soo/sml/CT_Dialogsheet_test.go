package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DialogsheetConstructor(t *testing.T) {
	v := sml.NewCT_Dialogsheet()
	if v == nil {
		t.Errorf("sml.NewCT_Dialogsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Dialogsheet should validate: %s", err)
	}
}

func TestCT_DialogsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Dialogsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Dialogsheet()
	xml.Unmarshal(buf, v2)
}
