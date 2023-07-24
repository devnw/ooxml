package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestRFC3066Constructor(t *testing.T) {
	v := terms.NewRFC3066()
	if v == nil {
		t.Errorf("terms.NewRFC3066 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.RFC3066 should validate: %s", err)
	}
}

func TestRFC3066MarshalUnmarshal(t *testing.T) {
	v := terms.NewRFC3066()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewRFC3066()
	xml.Unmarshal(buf, v2)
}
