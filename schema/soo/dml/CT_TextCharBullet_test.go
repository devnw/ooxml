package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextCharBulletConstructor(t *testing.T) {
	v := dml.NewCT_TextCharBullet()
	if v == nil {
		t.Errorf("dml.NewCT_TextCharBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextCharBullet should validate: %s", err)
	}
}

func TestCT_TextCharBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextCharBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextCharBullet()
	xml.Unmarshal(buf, v2)
}
