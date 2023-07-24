package powerpoint_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
)

func TestCT_RelConstructor(t *testing.T) {
	v := powerpoint.NewCT_Rel()
	if v == nil {
		t.Errorf("powerpoint.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed powerpoint.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := powerpoint.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := powerpoint.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
