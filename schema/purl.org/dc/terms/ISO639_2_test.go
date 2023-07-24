package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestISO639_2Constructor(t *testing.T) {
	v := terms.NewISO639_2()
	if v == nil {
		t.Errorf("terms.NewISO639_2 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ISO639_2 should validate: %s", err)
	}
}

func TestISO639_2MarshalUnmarshal(t *testing.T) {
	v := terms.NewISO639_2()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewISO639_2()
	xml.Unmarshal(buf, v2)
}
