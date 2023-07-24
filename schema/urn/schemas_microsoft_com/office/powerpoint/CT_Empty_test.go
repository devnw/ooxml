package powerpoint_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
)

func TestCT_EmptyConstructor(t *testing.T) {
	v := powerpoint.NewCT_Empty()
	if v == nil {
		t.Errorf("powerpoint.NewCT_Empty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed powerpoint.CT_Empty should validate: %s", err)
	}
}

func TestCT_EmptyMarshalUnmarshal(t *testing.T) {
	v := powerpoint.NewCT_Empty()
	buf, _ := xml.Marshal(v)
	v2 := powerpoint.NewCT_Empty()
	xml.Unmarshal(buf, v2)
}
