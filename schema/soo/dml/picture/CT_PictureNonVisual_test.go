package picture_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/picture"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := picture.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("picture.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed picture.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := picture.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := picture.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
