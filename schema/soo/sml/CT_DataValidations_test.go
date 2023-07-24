package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataValidationsConstructor(t *testing.T) {
	v := sml.NewCT_DataValidations()
	if v == nil {
		t.Errorf("sml.NewCT_DataValidations must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataValidations should validate: %s", err)
	}
}

func TestCT_DataValidationsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataValidations()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataValidations()
	xml.Unmarshal(buf, v2)
}
