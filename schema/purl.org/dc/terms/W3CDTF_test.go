package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestW3CDTFConstructor(t *testing.T) {
	v := terms.NewW3CDTF()
	if v == nil {
		t.Errorf("terms.NewW3CDTF must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.W3CDTF should validate: %s", err)
	}
}

func TestW3CDTFMarshalUnmarshal(t *testing.T) {
	v := terms.NewW3CDTF()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewW3CDTF()
	xml.Unmarshal(buf, v2)
}
