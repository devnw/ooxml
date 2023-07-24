package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextBulletTypefaceConstructor(t *testing.T) {
	v := dml.NewEG_TextBulletTypeface()
	if v == nil {
		t.Errorf("dml.NewEG_TextBulletTypeface must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBulletTypeface should validate: %s", err)
	}
}

func TestEG_TextBulletTypefaceMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBulletTypeface()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBulletTypeface()
	xml.Unmarshal(buf, v2)
}
