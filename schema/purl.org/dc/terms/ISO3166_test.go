package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestISO3166Constructor(t *testing.T) {
	v := terms.NewISO3166()
	if v == nil {
		t.Errorf("terms.NewISO3166 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ISO3166 should validate: %s", err)
	}
}

func TestISO3166MarshalUnmarshal(t *testing.T) {
	v := terms.NewISO3166()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewISO3166()
	xml.Unmarshal(buf, v2)
}
