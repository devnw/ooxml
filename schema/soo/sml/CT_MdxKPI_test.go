package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MdxKPIConstructor(t *testing.T) {
	v := sml.NewCT_MdxKPI()
	if v == nil {
		t.Errorf("sml.NewCT_MdxKPI must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MdxKPI should validate: %s", err)
	}
}

func TestCT_MdxKPIMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MdxKPI()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MdxKPI()
	xml.Unmarshal(buf, v2)
}
