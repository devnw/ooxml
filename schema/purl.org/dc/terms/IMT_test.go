package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestIMTConstructor(t *testing.T) {
	v := terms.NewIMT()
	if v == nil {
		t.Errorf("terms.NewIMT must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.IMT should validate: %s", err)
	}
}

func TestIMTMarshalUnmarshal(t *testing.T) {
	v := terms.NewIMT()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewIMT()
	xml.Unmarshal(buf, v2)
}
