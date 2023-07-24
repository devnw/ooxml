package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ShowInfoBrowseConstructor(t *testing.T) {
	v := pml.NewCT_ShowInfoBrowse()
	if v == nil {
		t.Errorf("pml.NewCT_ShowInfoBrowse must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ShowInfoBrowse should validate: %s", err)
	}
}

func TestCT_ShowInfoBrowseMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ShowInfoBrowse()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ShowInfoBrowse()
	xml.Unmarshal(buf, v2)
}
