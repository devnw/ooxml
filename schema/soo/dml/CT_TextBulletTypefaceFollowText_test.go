package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBulletTypefaceFollowTextConstructor(t *testing.T) {
	v := dml.NewCT_TextBulletTypefaceFollowText()
	if v == nil {
		t.Errorf("dml.NewCT_TextBulletTypefaceFollowText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBulletTypefaceFollowText should validate: %s", err)
	}
}

func TestCT_TextBulletTypefaceFollowTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBulletTypefaceFollowText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBulletTypefaceFollowText()
	xml.Unmarshal(buf, v2)
}
