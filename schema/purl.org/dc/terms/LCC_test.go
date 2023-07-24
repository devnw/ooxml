package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestLCCConstructor(t *testing.T) {
	v := terms.NewLCC()
	if v == nil {
		t.Errorf("terms.NewLCC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.LCC should validate: %s", err)
	}
}

func TestLCCMarshalUnmarshal(t *testing.T) {
	v := terms.NewLCC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewLCC()
	xml.Unmarshal(buf, v2)
}
