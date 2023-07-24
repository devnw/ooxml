package powerpoint_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
)

func TestTextdataConstructor(t *testing.T) {
	v := powerpoint.NewTextdata()
	if v == nil {
		t.Errorf("powerpoint.NewTextdata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed powerpoint.Textdata should validate: %s", err)
	}
}

func TestTextdataMarshalUnmarshal(t *testing.T) {
	v := powerpoint.NewTextdata()
	buf, _ := xml.Marshal(v)
	v2 := powerpoint.NewTextdata()
	xml.Unmarshal(buf, v2)
}
