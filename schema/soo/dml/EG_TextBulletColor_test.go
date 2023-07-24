package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextBulletColorConstructor(t *testing.T) {
	v := dml.NewEG_TextBulletColor()
	if v == nil {
		t.Errorf("dml.NewEG_TextBulletColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBulletColor should validate: %s", err)
	}
}

func TestEG_TextBulletColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBulletColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBulletColor()
	xml.Unmarshal(buf, v2)
}
