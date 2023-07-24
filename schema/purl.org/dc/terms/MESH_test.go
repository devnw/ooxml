package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestMESHConstructor(t *testing.T) {
	v := terms.NewMESH()
	if v == nil {
		t.Errorf("terms.NewMESH must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.MESH should validate: %s", err)
	}
}

func TestMESHMarshalUnmarshal(t *testing.T) {
	v := terms.NewMESH()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewMESH()
	xml.Unmarshal(buf, v2)
}
