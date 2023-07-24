package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := pml.NewCT_Picture()
	if v == nil {
		t.Errorf("pml.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
