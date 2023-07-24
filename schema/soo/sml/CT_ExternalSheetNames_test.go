package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalSheetNamesConstructor(t *testing.T) {
	v := sml.NewCT_ExternalSheetNames()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalSheetNames must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalSheetNames should validate: %s", err)
	}
}

func TestCT_ExternalSheetNamesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalSheetNames()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalSheetNames()
	xml.Unmarshal(buf, v2)
}
