package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PCDKPIConstructor(t *testing.T) {
	v := sml.NewCT_PCDKPI()
	if v == nil {
		t.Errorf("sml.NewCT_PCDKPI must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PCDKPI should validate: %s", err)
	}
}

func TestCT_PCDKPIMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PCDKPI()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PCDKPI()
	xml.Unmarshal(buf, v2)
}
