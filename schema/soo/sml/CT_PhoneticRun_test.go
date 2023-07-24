package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PhoneticRunConstructor(t *testing.T) {
	v := sml.NewCT_PhoneticRun()
	if v == nil {
		t.Errorf("sml.NewCT_PhoneticRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PhoneticRun should validate: %s", err)
	}
}

func TestCT_PhoneticRunMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PhoneticRun()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PhoneticRun()
	xml.Unmarshal(buf, v2)
}
