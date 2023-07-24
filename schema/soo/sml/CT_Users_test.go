package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_UsersConstructor(t *testing.T) {
	v := sml.NewCT_Users()
	if v == nil {
		t.Errorf("sml.NewCT_Users must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Users should validate: %s", err)
	}
}

func TestCT_UsersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Users()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Users()
	xml.Unmarshal(buf, v2)
}
