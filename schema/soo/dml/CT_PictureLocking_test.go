package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PictureLockingConstructor(t *testing.T) {
	v := dml.NewCT_PictureLocking()
	if v == nil {
		t.Errorf("dml.NewCT_PictureLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PictureLocking should validate: %s", err)
	}
}

func TestCT_PictureLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PictureLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PictureLocking()
	xml.Unmarshal(buf, v2)
}
