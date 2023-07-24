package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestLCSHConstructor(t *testing.T) {
	v := terms.NewLCSH()
	if v == nil {
		t.Errorf("terms.NewLCSH must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.LCSH should validate: %s", err)
	}
}

func TestLCSHMarshalUnmarshal(t *testing.T) {
	v := terms.NewLCSH()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewLCSH()
	xml.Unmarshal(buf, v2)
}
