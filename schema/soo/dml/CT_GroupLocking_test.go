package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GroupLockingConstructor(t *testing.T) {
	v := dml.NewCT_GroupLocking()
	if v == nil {
		t.Errorf("dml.NewCT_GroupLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GroupLocking should validate: %s", err)
	}
}

func TestCT_GroupLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GroupLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GroupLocking()
	xml.Unmarshal(buf, v2)
}
