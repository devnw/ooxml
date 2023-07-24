package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataValidationConstructor(t *testing.T) {
	v := sml.NewCT_DataValidation()
	if v == nil {
		t.Errorf("sml.NewCT_DataValidation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataValidation should validate: %s", err)
	}
}

func TestCT_DataValidationMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataValidation()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataValidation()
	xml.Unmarshal(buf, v2)
}
