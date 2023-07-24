package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBulletColorFollowTextConstructor(t *testing.T) {
	v := dml.NewCT_TextBulletColorFollowText()
	if v == nil {
		t.Errorf("dml.NewCT_TextBulletColorFollowText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBulletColorFollowText should validate: %s", err)
	}
}

func TestCT_TextBulletColorFollowTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBulletColorFollowText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBulletColorFollowText()
	xml.Unmarshal(buf, v2)
}
