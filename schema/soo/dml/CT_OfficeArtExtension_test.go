package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_OfficeArtExtensionConstructor(t *testing.T) {
	v := dml.NewCT_OfficeArtExtension()
	if v == nil {
		t.Errorf("dml.NewCT_OfficeArtExtension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_OfficeArtExtension should validate: %s", err)
	}
}

func TestCT_OfficeArtExtensionMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_OfficeArtExtension()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_OfficeArtExtension()
	xml.Unmarshal(buf, v2)
}
