package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := wml.NewCT_Picture()
	if v == nil {
		t.Errorf("wml.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
