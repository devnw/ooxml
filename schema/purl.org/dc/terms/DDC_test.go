package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestDDCConstructor(t *testing.T) {
	v := terms.NewDDC()
	if v == nil {
		t.Errorf("terms.NewDDC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.DDC should validate: %s", err)
	}
}

func TestDDCMarshalUnmarshal(t *testing.T) {
	v := terms.NewDDC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewDDC()
	xml.Unmarshal(buf, v2)
}
