package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlPictureConstructor(t *testing.T) {
	v := dml.NewCT_GvmlPicture()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlPicture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlPicture should validate: %s", err)
	}
}

func TestCT_GvmlPictureMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlPicture()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlPicture()
	xml.Unmarshal(buf, v2)
}
