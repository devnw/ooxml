package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestRFC1766Constructor(t *testing.T) {
	v := terms.NewRFC1766()
	if v == nil {
		t.Errorf("terms.NewRFC1766 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.RFC1766 should validate: %s", err)
	}
}

func TestRFC1766MarshalUnmarshal(t *testing.T) {
	v := terms.NewRFC1766()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewRFC1766()
	xml.Unmarshal(buf, v2)
}
