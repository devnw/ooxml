package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TagsDataConstructor(t *testing.T) {
	v := pml.NewCT_TagsData()
	if v == nil {
		t.Errorf("pml.NewCT_TagsData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TagsData should validate: %s", err)
	}
}

func TestCT_TagsDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TagsData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TagsData()
	xml.Unmarshal(buf, v2)
}
