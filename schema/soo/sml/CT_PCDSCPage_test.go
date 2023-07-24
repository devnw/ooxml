package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PCDSCPageConstructor(t *testing.T) {
	v := sml.NewCT_PCDSCPage()
	if v == nil {
		t.Errorf("sml.NewCT_PCDSCPage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PCDSCPage should validate: %s", err)
	}
}

func TestCT_PCDSCPageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PCDSCPage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PCDSCPage()
	xml.Unmarshal(buf, v2)
}
