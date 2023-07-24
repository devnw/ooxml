package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PhotoAlbumConstructor(t *testing.T) {
	v := pml.NewCT_PhotoAlbum()
	if v == nil {
		t.Errorf("pml.NewCT_PhotoAlbum must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_PhotoAlbum should validate: %s", err)
	}
}

func TestCT_PhotoAlbumMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_PhotoAlbum()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_PhotoAlbum()
	xml.Unmarshal(buf, v2)
}
