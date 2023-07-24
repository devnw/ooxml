package powerpoint_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
)

func TestIscommentConstructor(t *testing.T) {
	v := powerpoint.NewIscomment()
	if v == nil {
		t.Errorf("powerpoint.NewIscomment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed powerpoint.Iscomment should validate: %s", err)
	}
}

func TestIscommentMarshalUnmarshal(t *testing.T) {
	v := powerpoint.NewIscomment()
	buf, _ := xml.Marshal(v)
	v2 := powerpoint.NewIscomment()
	xml.Unmarshal(buf, v2)
}
