package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumPicBulletConstructor(t *testing.T) {
	v := wml.NewCT_NumPicBullet()
	if v == nil {
		t.Errorf("wml.NewCT_NumPicBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumPicBullet should validate: %s", err)
	}
}

func TestCT_NumPicBulletMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumPicBullet()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumPicBullet()
	xml.Unmarshal(buf, v2)
}
