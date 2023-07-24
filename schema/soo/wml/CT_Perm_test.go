package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PermConstructor(t *testing.T) {
	v := wml.NewCT_Perm()
	if v == nil {
		t.Errorf("wml.NewCT_Perm must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Perm should validate: %s", err)
	}
}

func TestCT_PermMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Perm()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Perm()
	xml.Unmarshal(buf, v2)
}
