package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextBulletConstructor(t *testing.T) {
	v := dml.NewEG_TextBullet()
	if v == nil {
		t.Errorf("dml.NewEG_TextBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBullet should validate: %s", err)
	}
}

func TestEG_TextBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBullet()
	xml.Unmarshal(buf, v2)
}
