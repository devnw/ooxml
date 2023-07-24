package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ContentPartLockingConstructor(t *testing.T) {
	v := dml.NewCT_ContentPartLocking()
	if v == nil {
		t.Errorf("dml.NewCT_ContentPartLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ContentPartLocking should validate: %s", err)
	}
}

func TestCT_ContentPartLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ContentPartLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ContentPartLocking()
	xml.Unmarshal(buf, v2)
}
