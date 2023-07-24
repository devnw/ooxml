package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestAG_LockingConstructor(t *testing.T) {
	v := dml.NewAG_Locking()
	if v == nil {
		t.Errorf("dml.NewAG_Locking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.AG_Locking should validate: %s", err)
	}
}

func TestAG_LockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewAG_Locking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewAG_Locking()
	xml.Unmarshal(buf, v2)
}
