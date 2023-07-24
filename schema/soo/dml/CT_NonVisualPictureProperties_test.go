package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualPicturePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualPictureProperties()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualPictureProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualPictureProperties should validate: %s", err)
	}
}

func TestCT_NonVisualPicturePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualPictureProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualPictureProperties()
	xml.Unmarshal(buf, v2)
}
