package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LsdExceptionConstructor(t *testing.T) {
	v := wml.NewCT_LsdException()
	if v == nil {
		t.Errorf("wml.NewCT_LsdException must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LsdException should validate: %s", err)
	}
}

func TestCT_LsdExceptionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LsdException()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LsdException()
	xml.Unmarshal(buf, v2)
}
