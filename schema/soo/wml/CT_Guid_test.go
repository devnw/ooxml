package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_GuidConstructor(t *testing.T) {
	v := wml.NewCT_Guid()
	if v == nil {
		t.Errorf("wml.NewCT_Guid must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Guid should validate: %s", err)
	}
}

func TestCT_GuidMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Guid()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Guid()
	xml.Unmarshal(buf, v2)
}
