package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextAutonumberBulletConstructor(t *testing.T) {
	v := dml.NewCT_TextAutonumberBullet()
	if v == nil {
		t.Errorf("dml.NewCT_TextAutonumberBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextAutonumberBullet should validate: %s", err)
	}
}

func TestCT_TextAutonumberBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextAutonumberBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextAutonumberBullet()
	xml.Unmarshal(buf, v2)
}
