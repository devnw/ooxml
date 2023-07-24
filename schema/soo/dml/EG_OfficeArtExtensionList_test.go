package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_OfficeArtExtensionListConstructor(t *testing.T) {
	v := dml.NewEG_OfficeArtExtensionList()
	if v == nil {
		t.Errorf("dml.NewEG_OfficeArtExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_OfficeArtExtensionList should validate: %s", err)
	}
}

func TestEG_OfficeArtExtensionListMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_OfficeArtExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_OfficeArtExtensionList()
	xml.Unmarshal(buf, v2)
}
