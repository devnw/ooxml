package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextBulletSizeConstructor(t *testing.T) {
	v := dml.NewEG_TextBulletSize()
	if v == nil {
		t.Errorf("dml.NewEG_TextBulletSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBulletSize should validate: %s", err)
	}
}

func TestEG_TextBulletSizeMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBulletSize()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBulletSize()
	xml.Unmarshal(buf, v2)
}
