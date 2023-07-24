package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_OfficeArtExtensionListConstructor(t *testing.T) {
	v := dml.NewCT_OfficeArtExtensionList()
	if v == nil {
		t.Errorf("dml.NewCT_OfficeArtExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_OfficeArtExtensionList should validate: %s", err)
	}
}

func TestCT_OfficeArtExtensionListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_OfficeArtExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_OfficeArtExtensionList()
	xml.Unmarshal(buf, v2)
}
