package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBulletSizeFollowTextConstructor(t *testing.T) {
	v := dml.NewCT_TextBulletSizeFollowText()
	if v == nil {
		t.Errorf("dml.NewCT_TextBulletSizeFollowText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBulletSizeFollowText should validate: %s", err)
	}
}

func TestCT_TextBulletSizeFollowTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBulletSizeFollowText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBulletSizeFollowText()
	xml.Unmarshal(buf, v2)
}
