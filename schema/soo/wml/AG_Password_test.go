package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestAG_PasswordConstructor(t *testing.T) {
	v := wml.NewAG_Password()
	if v == nil {
		t.Errorf("wml.NewAG_Password must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.AG_Password should validate: %s", err)
	}
}

func TestAG_PasswordMarshalUnmarshal(t *testing.T) {
	v := wml.NewAG_Password()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewAG_Password()
	xml.Unmarshal(buf, v2)
}
