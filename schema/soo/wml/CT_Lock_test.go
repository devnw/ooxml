package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LockConstructor(t *testing.T) {
	v := wml.NewCT_Lock()
	if v == nil {
		t.Errorf("wml.NewCT_Lock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Lock should validate: %s", err)
	}
}

func TestCT_LockMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Lock()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Lock()
	xml.Unmarshal(buf, v2)
}
