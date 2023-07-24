package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestBoxConstructor(t *testing.T) {
	v := terms.NewBox()
	if v == nil {
		t.Errorf("terms.NewBox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Box should validate: %s", err)
	}
}

func TestBoxMarshalUnmarshal(t *testing.T) {
	v := terms.NewBox()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewBox()
	xml.Unmarshal(buf, v2)
}
