package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestPointConstructor(t *testing.T) {
	v := terms.NewPoint()
	if v == nil {
		t.Errorf("terms.NewPoint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Point should validate: %s", err)
	}
}

func TestPointMarshalUnmarshal(t *testing.T) {
	v := terms.NewPoint()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewPoint()
	xml.Unmarshal(buf, v2)
}
