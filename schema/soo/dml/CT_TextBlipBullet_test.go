package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBlipBulletConstructor(t *testing.T) {
	v := dml.NewCT_TextBlipBullet()
	if v == nil {
		t.Errorf("dml.NewCT_TextBlipBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBlipBullet should validate: %s", err)
	}
}

func TestCT_TextBlipBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBlipBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBlipBullet()
	xml.Unmarshal(buf, v2)
}
