package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestTGNConstructor(t *testing.T) {
	v := terms.NewTGN()
	if v == nil {
		t.Errorf("terms.NewTGN must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.TGN should validate: %s", err)
	}
}

func TestTGNMarshalUnmarshal(t *testing.T) {
	v := terms.NewTGN()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewTGN()
	xml.Unmarshal(buf, v2)
}
