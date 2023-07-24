package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartGalleryConstructor(t *testing.T) {
	v := wml.NewCT_DocPartGallery()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartGallery must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartGallery should validate: %s", err)
	}
}

func TestCT_DocPartGalleryMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartGallery()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartGallery()
	xml.Unmarshal(buf, v2)
}
