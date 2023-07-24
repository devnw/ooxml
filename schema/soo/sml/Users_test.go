package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestUsersConstructor(t *testing.T) {
	v := sml.NewUsers()
	if v == nil {
		t.Errorf("sml.NewUsers must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Users should validate: %s", err)
	}
}

func TestUsersMarshalUnmarshal(t *testing.T) {
	v := sml.NewUsers()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewUsers()
	xml.Unmarshal(buf, v2)
}
