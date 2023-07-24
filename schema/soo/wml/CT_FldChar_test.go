package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FldCharConstructor(t *testing.T) {
	v := wml.NewCT_FldChar()
	if v == nil {
		t.Errorf("wml.NewCT_FldChar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FldChar should validate: %s", err)
	}
}

func TestCT_FldCharMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FldChar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FldChar()
	xml.Unmarshal(buf, v2)
}
