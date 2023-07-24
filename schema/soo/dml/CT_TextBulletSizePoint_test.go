package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBulletSizePointConstructor(t *testing.T) {
	v := dml.NewCT_TextBulletSizePoint()
	if v == nil {
		t.Errorf("dml.NewCT_TextBulletSizePoint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBulletSizePoint should validate: %s", err)
	}
}

func TestCT_TextBulletSizePointMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBulletSizePoint()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBulletSizePoint()
	xml.Unmarshal(buf, v2)
}
